package main

import (
	"mitra/controller"
	"mitra/controller/admin"
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

	r.Post("/attendees", func(w http.ResponseWriter, r *http.Request) {})

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

	r.Route("/attendees", func(r chi.Router) {
		r.Get("/", admin.ListAttendees)
		r.Get("/{id}", admin.GetAttendees)
		r.Delete("/{id}", admin.DeleteAttendees)
	})

	r.Route("/log", func(r chi.Router) {
		r.Get("/dwell", admin.ListDwellTimeLogs)
		r.Get("/dwell/{user_id}", admin.ListDwellTimeLogs)
		r.Get("/click", admin.ListClickLogs)
		r.Get("/click/{user_id}", admin.ListClickLogs)
	})

	return r
}
