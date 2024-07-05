package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r) 
		return
	}
	w.Write([]byte("Display the home page"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	
	id := r.URL.Query().Get("id")

	Id, err := strconv.Atoi(id)
	if err != nil || Id < 1 {
		http.Error(w, "404 page not found", 404)
		return
	}

	info := fmt.Sprintf("number of snipet id = %s", id)
	w.Write([]byte(info + "\n"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		file, err := os.Open("index.html")
		if err != nil {
			fmt.Println(false)
			return
		}
		dat, err := ioutil.ReadAll(file)
		fmt.Println(string(dat))
		if err != nil {
			fmt.Println(true)
			return
		}
		fmt.Println(http.DetectContentType([]byte(dat))) // для хедера тип контента

		w.Header()["anime-type"] = []string{"1;mode=block"}
		w.Header().Set("cache-control", "naruto uzumaki")
		// In contrast, the Add() method appends a new "Cache-Control" header and can
		// be called multiple times.
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")
		w.Header()["Date"] = nil

		// Delete all values for the "Cache-Control" header.
		// Retrieve the first value for the "Cache-Control" header.

		// дефолтные хедеры не удалить таким образам
		data := w.Header().Get("Cache-Control")
		fmt.Println(data)
		// w.Header().Set("Allow", "POST")
		// w.Header().Add("Date", "baribir") // что бы изменить значения уже готового хедера
		// content := w.Header().Get("Date") // получить значения хедера

		// fmt.Println(content)
		// w.Header().Del("Allow")
		http.Error(w, "Method Not Allowed", 405) // w.WriteHeader(405) => использовать тока один раз && w.Write([]byte("some text"))
		return
	}

	fmt.Fprintln(w, "Create a new snippet")
}

// PAGE 41
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("starting server in http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}