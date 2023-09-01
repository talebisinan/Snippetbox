package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"snippetbox.sinantalebi.net/internal/models"
	"snippetbox.sinantalebi.net/ui"
)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to the HTML templates.
type TemplateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
}

func HumanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"HumanDate": HumanDate,
}

// Instead of reading from disk every time a page is requested, it only reads
// the templates once at the start of the application.
func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		paterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		// Parse the base template file into a template set and use the Funcs() method to register functions.
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, paterns...)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map as normal...
		cache[name] = ts
	}
	return cache, nil
}
