package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippet)
	log.Println("starting server in http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
