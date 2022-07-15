package main

import (
	"mitra/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Keep greeting to world forever"))
	})

	r.Route("/search", func(r chi.Router) {
		// ListSearchPage
		r.Get("/", controller.ListSearchResult)
	})

	r.Route("/log", func(r chi.Router) {
		r.Post("/dwell", controller.CreateDwellTimeLog)
		r.Post("/click", controller.CreateClickLog)
	})

	return r
}

func adminRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(adminOnly)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})

	r.Route("/log", func(r chi.Router) {
		r.Get("/dwell", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/click", func(w http.ResponseWriter, r *http.Request) {})
	})

	return r
}
