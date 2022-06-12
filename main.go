package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Book struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publicationYear"`
	Genre           string `json:"genre"`
}

func retrieveBooks() []Book {
	fileData, _ := os.ReadFile("data.json")
	var books []Book
	json.Unmarshal(fileData, &books)

	return books
}

func saveBook(book Book) error {
	books := retrieveBooks()
	books = append(books, book)
	data, err := json.Marshal(books)

	if err != nil {
		return err
	}

	os.WriteFile("data.json", data, 0644)

	return nil
}

func ListBooks(w http.ResponseWriter) {
	books := retrieveBooks()
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	body, readErr := ioutil.ReadAll(r.Body)

	if readErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to read body"))
		return
	}

	var book Book
	unmarshalErr := json.Unmarshal(body, &book)

	if unmarshalErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body must contain id, title, author, publicationYear and genre"))
		return
	}

	// save book to "database"
	saveErr := saveBook(book)

	if saveErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to save book"))
		return
	}

	// respond with the newly created book
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func HandleBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		ListBooks(w)
	case "POST":
		CreateBook(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Method " + r.Method + " not allowed"))
	}
}

func main() {
	http.HandleFunc("/books", HandleBooks)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
