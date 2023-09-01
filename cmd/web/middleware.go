package main

import (
	"fmt"
	"net/http"
)

func (app *Application) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func (app *Application) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use the built-in recover function to check if there has been a
		// panic or not.
		defer func() {
			if err := recover(); err != nil {
				// Set a "Connection: close" header on the response.
				w.Header().Set("Connection", "close")

				// Call the serverError helper method to return a 500 Internal
				// Server response.
				app.serverError(w, fmt.Errorf("%s", err))

			}
		}()

		next.ServeHTTP(w, r)
	})
}

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// https://owasp.org/www-project-secure-headers/
		// Set the X-XSS-Protection header so that the browser stops pages from loading
		// when they detect reflected cross-site scripting (XSS) attacks.
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")

		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")

		// Instructs browsers to not MIME-type sniff the content-type of the response
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Deny is used to help prevent clickjacking attacks in older browsers that don’t support CSP headers
		w.Header().Set("X-Frame-Options", "deny")

		// Disable the blocking of cross-site scripting attacks
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

func (app *Application) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.IsAuthenticated(r) {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
