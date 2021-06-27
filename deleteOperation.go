package main

import (
	"encoding/json"
	"fmt"

	"io/ioutill"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}

var Articles []Article

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page ")
	fmt.Println("End point Hit : Homepage ")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point Hit : returnAllArticles ")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Id"]
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
 	// return the string response containing the request body
 	reqBody := ioutill.ReadAll(r.Body)
 	fmt.Fprintf(w, "%+v", string(reqBody))
 	var article Article
 	json.Unmarshal(reqBody, &article)
 	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func delteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Id"]
	for index, article := range Articles {
		if article.Id == key {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Homepage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{Id}", returnSingleArticle)
	// we’ll be adding .Methods("POST") to the end of our
	// route to specify that we only want to call this
	// function when the incoming request is a HTTP POST reques
	// myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("article/{id}", delteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9970", myRouter))
}
func main() {
	Articles = []Article{
		{Id: "1", Title: "Article 1", Desc: "Article 1 Description", Content: "Articles 1 Content"},
		{Id: "2", Title: "Article 2", Desc: "Article 2 Description", Content: "Articles 2 Content"},
		{Id: "3", Title: "Article 3", Desc: "Article 3 Description", Content: "Articles 3 Content"},
	}

	handleRequest()
}
