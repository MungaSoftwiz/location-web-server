package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MungaSoftwiz/location-web-server/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", handlers.HelloHandler)

	log.Println("Starting server on :8080")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe("0.0.0.0:"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
