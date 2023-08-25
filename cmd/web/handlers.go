package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/julienschmidt/httprouter"
	"snippetbox.sinantalebi.net/internal/models"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	templateData := app.NewTemplateData(r)
	templateData.Snippets = snippets

	app.renderPage(w, http.StatusOK, "home.tmpl", templateData)

}

func (app *Application) snippetView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	templateData := app.NewTemplateData(r)
	templateData.Snippet = snippet

	app.renderPage(w, http.StatusOK, "view.tmpl", templateData)
}

type SnippetCreateForm struct {
	Title       string
	Content     string
	Expires     string
	FieldErrors map[string]string
}

func (app *Application) showSnippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.NewTemplateData(r)
	data.Form = SnippetCreateForm{
		Expires: "365",
	}
	app.renderPage(w, http.StatusOK, "create.tmpl", data)
}

func (app *Application) doSnippetCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := SnippetCreateForm{
		Title:       r.PostForm.Get("title"),
		Content:     r.PostForm.Get("content"),
		Expires:     r.PostForm.Get("expires"),
		FieldErrors: map[string]string{},
	}

	if strings.TrimSpace(form.Title) == "" {
		form.FieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(form.Title) > 100 {
		form.FieldErrors["title"] = "This field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(form.Content) == "" {
		form.FieldErrors["content"] = "This field cannot be blank"
	}

	if form.Expires != "1" && form.Expires != "7" && form.Expires != "365" {
		form.FieldErrors["expires"] = "This field is invalid"
	}

	if len(form.FieldErrors) > 0 {
		data := app.NewTemplateData(r)
		data.Form = form
		app.renderPage(w, http.StatusBadRequest, "create.tmpl", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
