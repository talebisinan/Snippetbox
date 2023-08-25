package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	router.HandlerFunc(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileserver).ServeHTTP)

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.showSnippetCreate)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.doSnippetCreate)

	standard := alice.New(app.RecoverPanic, app.LogRequest, SecureHeaders)

	return standard.Then(router)
}
