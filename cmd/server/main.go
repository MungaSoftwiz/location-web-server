package main

import (
	"log"
	"net/http"

	"github.com/MungaSoftwiz/location-web-server/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handlers.HelloHandler)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}