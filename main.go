package main

import (
	"encoding/json"
	"log"
	"net/http"

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
	w.Header().Set("content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

//Get single Books
func getBook(w http.ResponseWriter, r *http.Request) {

}

//crewate Book
func createBooks(w http.ResponseWriter, r *http.Request) {

}

//update book
func updateBooks(w http.ResponseWriter, r *http.Request) {

}

//delete book
func deleteBooks(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// init Router
	r := mux.NewRouter()

	//mock data
	books = append(books, Book{ID: "1", Isbn: "44334", Title: "Book one", Author: &Author{Firstname: "Bawantha", Lastname: "Thilan"}})
	books = append(books, Book{ID: "2" , Isbn: "44334", Title:"Book two" , Author: &Author{Firstname: "Joen" , Lastname: "Doe"} } )

	//Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
