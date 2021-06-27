package main

import (
	"fmt"
	"net/http"
)

func main() {

	//handling simple "/" route using handeler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Server User")
	})

	//handling "/teachers" route using handeler function
	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Server Teachers")
	})

	//handling "/students" route using handeler function
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to server students")
	})

	http.ListenAndServe(":7600", nil)
}
