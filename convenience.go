package main

import "net/http"

// These are convenience functions to essentially alias functionality

func (a *application) get(s string, h http.HandlerFunc) {
	a.App.Routes.Get(s, h)
}

func (a *application) post(s string, h http.HandlerFunc) {
	a.App.Routes.Post(s, h)
}

// for middleware
func (a *application) use(m ...func(http.Handler) http.Handler) {
	a.App.Routes.Use(m...)
}