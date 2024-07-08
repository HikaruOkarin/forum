package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/homepage.html",
		"./ui/html/base_layout.html",
		"./ui/html/footer_partial.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return

	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.Error(w, "404 page not found", 404)
		return
	}

	info := fmt.Sprintf("number of snipet id = %d", id)
	w.Write([]byte(info + "\n"))
}

func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.Header()["anime-type"] = []string{"1;mode=block"}
		w.Header().Set("cache-control", "naruto uzumaki")
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")
		w.Header()["Date"] = nil //for delete system generated headers

		data := w.Header().Get("Cache-Control")
		fmt.Println(data)

		w.Header().Del("Allow")                  // for delete header
		http.Error(w, "Method Not Allowed", 405) // w.WriteHeader(405) => использовать тока один раз && w.Write([]byte("some text"))
		return
	}

	fmt.Fprintln(w, "Create a new snippet")
}
