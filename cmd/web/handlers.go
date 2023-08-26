package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"snippetbox.sinantalebi.net/internal/models"
	"snippetbox.sinantalebi.net/internal/validator"
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
	Title   string `form:"title"`
	Content string `form:"content"`
	Expires string `form:"expires"`

	validator.Validator `form:"-"` // ignored
}

func (app *Application) doSnippetCreate(w http.ResponseWriter, r *http.Request) {
	var form SnippetCreateForm
	err := app.DecodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field is too long (maximum is 100 characters)")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedString(form.Expires, "1", "7", "365"), "expires", "This field is invalid")

	if !form.Valid() {
		data := app.NewTemplateData(r)
		data.Form = form
		app.renderPage(w, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *Application) showSnippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.NewTemplateData(r)
	data.Form = SnippetCreateForm{
		Expires: "365",
	}
	app.renderPage(w, http.StatusOK, "create.tmpl", data)
}
