package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

var books []Book

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("No book found with specified ID"))
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)

	books = append(books, newBook)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for index, book := range books {
		if book.ID == id {
			var updatedBook Book
			json.NewDecoder(r.Body).Decode(&updatedBook)
			books[index] = updatedBook

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("No book found with specified ID"))
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Book deleted successfully"))
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("No book found with specified ID"))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
