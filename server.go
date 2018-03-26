package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello) //handles "/" directory

	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
