package main

import (
	"log"
	"net/http"

	"backend/internal/db"
	"backend/req/http/router"
)

func main() {
	db.NewMySQL()

	mux := http.NewServeMux()

	router.RegisterRoutes(mux)

	log.Println("API listening on :5000")
	log.Fatal(http.ListenAndServe(":5000", mux))
}
