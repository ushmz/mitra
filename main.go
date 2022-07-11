package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
)

var gendoc = flag.Bool("docgen", false, "Generate router documentation")

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/api", router())
	r.Mount("/admin", adminRouter())

	if *gendoc {
		fmt.Println(docgen.JSONRoutesDoc(r))
	}

	http.ListenAndServe(":3333", r)
}

func router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Keep greeting to world forever"))
	})

	r.Route("/search", func(r chi.Router) {
		// ListSearchPage
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	})

	r.Route("/log", func(r chi.Router) {
		r.Post("/dwell", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/click", func(w http.ResponseWriter, r *http.Request) {})
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
