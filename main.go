package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Struct(Model)
type Book struct {
	ID     string  `json:"id`
	Isbn   string  `json:"isbn`
	Title  string  `json:"title`
	Author *Author `json:"author`
}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//init books var as a slice book struct
var books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get single Books
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application.json")
	params := mux.Vars(r) //Get params
	//Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//crewate Book
func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //Mock ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

//update book
func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}

	}
	json.NewEncoder(w).Encode(books)
}

//delete book
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// init Router
	r := mux.NewRouter()

	//mock data
	books = append(books, Book{ID: "1", Isbn: "44334", Title: "Book one", Author: &Author{Firstname: "Bawantha", Lastname: "Thilan"}})
	books = append(books, Book{ID: "2", Isbn: "44334", Title: "Book two", Author: &Author{Firstname: "Joen", Lastname: "Doe"}})

	//Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
