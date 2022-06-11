package main

import (
	"encoding/json"
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

func HandleBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fileData, _ := os.ReadFile("data.json")
	var books []Book
	json.Unmarshal(fileData, &books)

	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/books", HandleBooks)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
