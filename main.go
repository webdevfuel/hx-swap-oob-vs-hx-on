package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/webdevfuel/hx-swap-oob-vs-hx-on/template"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := template.Index()
		component.Render(r.Context(), w)
	})
	r.Post("/items", func(w http.ResponseWriter, r *http.Request) {
		value := r.FormValue("item")
		component := template.FormWithLatestItem(value, true)
		component.Render(r.Context(), w)
	})
	http.ListenAndServe("localhost:3000", r)
}
