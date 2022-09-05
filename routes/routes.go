package routes

import (
	"encoding/json"
	"net/http"
	users "nickm980/user"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	users.RegisterUserRoutes(r)
	r.HandleFunc("*", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"statusCode": 404, "success": false, "message": "Page not found"})
	})
}
