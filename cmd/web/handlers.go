package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
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
		app.errorLog.Println(err.Error())

		http.Error(w, "Internal Server Error", 500)
		return

	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (app *application) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, "404 page not found", 404)
		return
	}
	info := fmt.Sprintf("number of snippet id = %d", id)
	w.Write([]byte(info + "\n"))
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new post"))
}
