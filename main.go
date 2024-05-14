package main

import (
	"log"
	"net/http"
	"os"

	"github.com/John-AG/go-rest-api/models"
	"github.com/John-AG/go-rest-api/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	models.InitDB(dbURL)

	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
