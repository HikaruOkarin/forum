package main

import (
	"fmt"
	"forum/pkg/models"

	"net/http"
	"strconv"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	s, err := app.posts.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{Posts: s}

	app.render(w, r, "home.page.tmpl", data)

}

func (app *application) ShowSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.posts.Get(id)
	data := &templateData{Post: s}
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.tmpl", data)

}

func (app *application) CreateSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)

}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {

		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")
	id, err := app.posts.Insert(title, content, expires)
	if err != nil {

		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)

}
