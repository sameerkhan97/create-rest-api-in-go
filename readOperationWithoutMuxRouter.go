package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Our Home Page")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBooks")
	json.NewEncoder(w).Encode(Books)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/Books", returnAllBooks)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	Books = []Book{
		{Title: "Book 1", Desc: "Boook 1 Description", Content: "Books 1 Content"},
		{Title: "Book 2", Desc: "Boook 2 Description", Content: "Books 2 Content"},
		{Title: "Book 3", Desc: "Boook 3 Description", Content: "Books 3 Content"},
	}
	handleRequest()
}
