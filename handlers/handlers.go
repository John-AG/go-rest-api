package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Welcome to my Go Rest Api"}
	json.NewEncoder(w).Encode(response)
}
