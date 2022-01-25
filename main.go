package main

import (
	"log"
	"net/http"
	"simple/handler"

	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	r := mux.NewRouter()

	// Endpoint
	r.HandleFunc("/api/books", handler.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", handler.GetBook).Methods("GET")
	r.HandleFunc("/api/books", handler.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", handler.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", handler.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
