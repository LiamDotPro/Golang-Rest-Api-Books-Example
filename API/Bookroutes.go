package API

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"math/rand"
)

// Book Structs (Model)

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var books []Book

func main() {

	// init the router
	router := mux.NewRouter()

	// Seed Data
	books = append(books, Book{ID: "1", Isbn: "qweqwe-1234", Title: "Book One", Author: &Author{FirstName: "Liam", LastName: "Read"}})

	books = append(books, Book{ID: "2", Isbn: "qweqwe-1234", Title: "Book Two", Author: &Author{FirstName: "Steve", LastName: "Read"}})

	// Router Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // Get any incoming params

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Book{})
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

func updateBooks(writer http.ResponseWriter, request *http.Request) {

}

func deleteBook(writer http.ResponseWriter, request *http.Request) {

}
