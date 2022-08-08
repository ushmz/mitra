package main

import (
	"mitra/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)

	h := handler.NewHandler()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Keep greeting to world forever"))
	})

	r.Route("/attendees", func(r chi.Router) {
		r.Post("/", h.User.CreateUser)
	})

	r.Route("/search", func(r chi.Router) {
		// ListSearchPage
		r.Get("/", h.Search.ListSearchResult)
	})

	r.Route("/log", func(r chi.Router) {
		r.Post("/dwell", h.Log.CreateDwellTimeLog)
		r.Post("/click", h.Log.CreateClickLog)
		r.Post("/hover", h.Log.CreateHoverLog)
	})

	return r
}

func adminRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(adminOnly)

	h := handler.NewAdminHandler()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.ListUsers)
		r.Get("/{id}", h.GetUser)
		r.Delete("/{id}", h.DeleteUser)
	})

	r.Route("/log", func(r chi.Router) {
		r.Get("/dwell", h.ListDwellTimeLogs)
		r.Get("/dwell/{user_id}", h.ListDwellTimeLogs)
		r.Get("/click", h.ListClickLogs)
		r.Get("/click/{user_id}", h.ListClickLogs)
		r.Get("/hover", h.ListHoverLogs)
		r.Get("/hover/{user_id}", h.ListHoverLogs)
	})

	return r
}
