package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	files := []string{
		"./ui/html/homepage.html",
		"./ui/html/base_layout.html",
		"./ui/html/footer_partial.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return

	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)

		return
	}
}

func (app *application) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	info := fmt.Sprintf("number of snippet id = %d", id)
	w.Write([]byte(info + "\n"))
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	w.Header().Set("Allow", "POST")
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n Kobayashi"
	expires := "7"
	id, err := app.posts.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
