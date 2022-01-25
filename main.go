package main

import (
	"log"
	"net/http"
	"simple/handler"
	"simple/model"

	"github.com/gorilla/mux"
)

var books []model.Book

func main() {
	// Init Router
	r := mux.NewRouter()

	// mock data
	books = append(books, model.Book{ID: "1", Isbn: "789456", Title: "First Book", Author: &model.Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, model.Book{ID: "2", Isbn: "564123", Title: "Second Book", Author: &model.Author{Firstname: "Steve", Lastname: "Smith"}})

	// Endpoint
	r.HandleFunc("/api/books", handler.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", handler.GetBook).Methods("GET")
	r.HandleFunc("/api/books", handler.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", handler.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", handler.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
