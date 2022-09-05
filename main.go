package main

import (
	"encoding/json"
	"log"
	"net/http"
	"nickm980/middleware"
	"nickm980/repository"
	"nickm980/routes"
	user "nickm980/user"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/pkg/profile"
)

var port string = "10000"

func main() {
	defer profile.Start(profile.MemProfile).Stop()
	go func() {
		http.ListenAndServe(":8081", nil)
	}()

	initDatabase()
	startMuxServer()

}

func initDatabase() {
	log.Default().Println("Connecting to database server")
	repository.Connect()
}

func startMuxServer() {
	log.Default().Println("Starting server")
	r := mux.NewRouter()
	api := r.PathPrefix("/api/").Subrouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"statusCode": 404, "success": false, "message": "The requested resource does not exist"})
	})

	r.StrictSlash(true)
	r.PathPrefix("/resources/").
		Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", renderTemplate)

	middleware.RegisterMiddleware(api)
	routes.RegisterRoutes(api)

	log.Default().Println("Server started. Listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {

	var u user.User = user.User{
		Name:     "j",
		Email:    "email",
		Password: "pass",
	}

	parsedTemplate, _ := template.ParseFiles("app/index.html")
	err := parsedTemplate.Execute(w, u)

	if err != nil {
		log.Fatal("Error executing template :", err)
		return
	}
}
