package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/John-AG/go-rest-api/models"
	"github.com/gorilla/mux"
)

var users = []models.User{
	{ID: "1", Name: "John Doe", Email: "JohnDoe@example.com"},
	{ID: "2", Name: "Jane Doe", Email: "JaneDoe@example.com"},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Welcome to my Go Rest Api"}
	json.NewEncoder(w).Encode(response)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}

type Response struct {
	Message string `json:"message"`
}
