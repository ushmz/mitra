package main

import (
	"mitra/handler"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func router(db *sqlx.DB, fb *firebase.App) chi.Router {
	r := chi.NewRouter()
	h := handler.NewHandler(db, fb)

	r.Post("/signup", h.User.CreateUser)

	return r
}

func authRouter(db *sqlx.DB, fb *firebase.App) chi.Router {
	r := chi.NewRouter()
	// r.Use(firebaseAuth(fb))

	h := handler.NewHandler(db, fb)

	r.Route("/user", func(r chi.Router) {
		r.Post("/assign", h.Task.AssignTask)
	})

	r.Route("/task", func(r chi.Router) {
		r.Get("/topics", h.Task.GetTaskQueries)
	})

	r.Route("/search", func(r chi.Router) {
		r.Get("/", h.Search.ListSearchResult)
		r.Get("/test", h.Search.GetSimilarweb)
		r.Get("/icon", h.Search.ListSearchResult)
		r.Get("/ratio", h.Search.ListSearchResult)
		r.Get("/purpose", h.Search.ListSearchResult)
	})

	r.Route("/log", func(r chi.Router) {
		r.Post("/dwell", h.Log.CreateDwellTimeLog)
		r.Post("/click", h.Log.CreateClickLog)
		r.Post("/hover", h.Log.CreateHoverLog)
	})

	return r
}

func adminRouter(fb *firebase.App) chi.Router {
	r := chi.NewRouter()
	// r.Use(firebaseAuth(fb))
	r.Use(adminOnly())

	h := handler.NewAdminHandler()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.User.ListUsers)
		r.Get("/{id}", h.User.GetUser)
		r.Delete("/{id}", h.User.DeleteUser)
	})

	r.Route("/log", func(r chi.Router) {
		r.Get("/dwell", h.Log.ListDwellTimeLogs)
		r.Get("/dwell/{user_id}", h.Log.ListDwellTimeLogs)
		r.Get("/click", h.Log.ListClickLogs)
		r.Get("/click/{user_id}", h.Log.ListClickLogs)
		r.Get("/hover", h.Log.ListHoverLogs)
		r.Get("/hover/{user_id}", h.Log.ListHoverLogs)
	})

	return r
}
