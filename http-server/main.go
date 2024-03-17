package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "sahith")
	})
	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "black")
	})
	http.HandleFunc("/reg", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212ad184")
	})
	//fmt.Printf("Server running (port=8080), route: http://localhost:8080/name")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
