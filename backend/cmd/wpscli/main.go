package main

import (
	"log"
	"net/http"
	"os"

	"backend/internal/db"
	"backend/req/http/router"
)

func main() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		panic("GEMINI_API_KEY is empty")
	}

	db.NewMySQL()

	mux := http.NewServeMux()

	router.RegisterRoutes(mux)

	log.Println("API listening on :5000")
	log.Fatal(http.ListenAndServe(":5000", mux))
}
