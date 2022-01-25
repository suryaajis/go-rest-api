package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"simple/model"
	"strconv"

	"github.com/gorilla/mux"
)

var books []model.Book

// func mockData() {
// 	books = append(books, model.Book{ID: "1", Isbn: "789456", Title: "First Book", Author: &model.Author{Firstname: "John", Lastname: "Doe"}})
// 	books = append(books, model.Book{ID: "2", Isbn: "564123", Title: "Second Book", Author: &model.Author{Firstname: "Steve", Lastname: "Smith"}})
// }

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// mockData()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// mockData()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100)) // Mock ID
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	filterBooks := []model.Book{}
	for _, book := range books {
		if book.ID != params["id"] {
			filterBooks = append(filterBooks, book)
		}
	}
	books = filterBooks

	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["id"]
	books = append(books, book)

	json.NewEncoder(w).Encode(books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	filterBooks := []model.Book{}
	params := mux.Vars(r) // Get Params
	for _, book := range books {
		if book.ID != params["id"] {
			filterBooks = append(filterBooks, book)
		}
	}

	books = filterBooks
	json.NewEncoder(w).Encode(books)
}
