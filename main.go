package main

import (
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

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {

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

	//Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
