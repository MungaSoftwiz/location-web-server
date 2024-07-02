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

	log.Println("Starting server on :4000")

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	err := http.ListenAndServe("0.0.0.0:"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
