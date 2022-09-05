package users

import (
	"encoding/json"
	"log"
	"net/http"
	errors "nickm980/utils"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/users/{name}", HandleGetUserById).Methods("GET")
	r.HandleFunc("/users", HandleCreateUser).Methods("POST")
}

func HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]

	if !ok {
		log.Default().Println("[HandleGetUserById] name variable failed to fetch")
		json.NewEncoder(w).Encode(map[string]interface{}{"statusCode": 400, "success": false, "message": name + " does not exist"})
		return
	}
	result, _ := FindUserByName(r.URL.Path)

	json.NewEncoder(w).Encode(result)
}

type createUserRequest struct {
	Name     string
	Email    string
	Password string
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var requestStruct createUserRequest
	json.NewDecoder(r.Body).Decode(&requestStruct)
	err := errors.Validate(w, requestStruct)

	if err != nil {
		return
	}

	_, creationStatus := CreateUser(requestStruct.Name, requestStruct.Email, requestStruct.Password)

	json.NewEncoder(w).Encode(creationStatus)
}
