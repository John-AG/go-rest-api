package main

import (
	"log"
	"net/http"

	"github.com/John-AG/go-rest-api/router"
)

func main() {
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
