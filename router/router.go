package router

import (
	"github.com/John-AG/go-rest-api/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	return r
}
