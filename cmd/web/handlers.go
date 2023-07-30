package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"ui/html/base.tmpl.html",
		"ui/html/partials/nav.tmpl.html",
		"ui/html/pages/home.tmpl.html",
	}

	templateSet, err := template.ParseFiles(files...)
	if err != nil {
		app.serveError(w, err)
		return
	}

	err = templateSet.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serveError(w, err)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

/*
* Custom handler that satisfies the http.Handler interface

	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}

*
*/
type customHandler struct{}

func (h *customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a custom handler!"))
}
