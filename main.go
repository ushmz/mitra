package main

import (
	"flag"
	"fmt"
	"mitra/auth"
	"mitra/config"
	"mitra/handler"
	"mitra/store"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
)

var gendoc = flag.Bool("docgen", false, "Generate router documentation")

func main() {
	err := config.Init()
	if err != nil {
		msg := fmt.Sprintf("Failed to load configurations\n%v", err)
		panic(msg)
	}

	db, err := store.InitDB()
	if err != nil {
		msg := fmt.Sprintf("Failed to connect DB\n%v", err)
		panic(msg)
	}

	fb, err := auth.InitFirebaseApp()
	if err != nil {
		msg := fmt.Sprintf("Failed to initialize firebase SDK\n%v", err)
		panic(msg)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(corsHandler())

	r.Get("/", handler.Index)

	r.Mount("/api/u", router(db, fb))
	r.Mount("/api/a", authRouter(db, fb))
	// r.Mount("/api/admin", adminRouter())

	flag.Parse()
	if *gendoc {
		fmt.Println(docgen.JSONRoutesDoc(r))
	}

	http.ListenAndServe(":3333", r)
}
