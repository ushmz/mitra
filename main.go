package main

import (
	"flag"
	"fmt"
	"mitra/handler"
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
		r.Get("/", handler.ListSearchResult)
	})

	r.Route("/log", func(r chi.Router) {
		r.Post("/dwell", handler.CreateDwellTimeLog)
		r.Post("/click", handler.CreateClickLog)
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

func adminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAdmin, ok := r.Context().Value("acl.admin").(bool)
		if !ok || !isAdmin {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
