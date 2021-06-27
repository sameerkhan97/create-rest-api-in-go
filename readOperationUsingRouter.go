package main

import (
	"encoding/json"
	"fmt"
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
	fmt.Fprintf(w, "Welcome to Our Home Page")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
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

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", Homepage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{Id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":2060", myRouter))

}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Article 1", Desc: "Article 1 Description", Content: "Articles 1 Content"},
		Article{Id: "2", Title: "Article 2", Desc: "Article 2 Description", Content: "Articles 2 Content"},
		Article{Id: "3", Title: "Article 3", Desc: "Article 3 Description", Content: "Articles 3 Content"},
	}

	handleRequest()
}
