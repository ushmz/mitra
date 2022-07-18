package main

import (
	"flag"
	"fmt"
	"mitra/config"
	"mitra/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
)

var gendoc = flag.Bool("docgen", false, "Generate router documentation")

func main() {
	if err := config.Init(); err != nil {
		msg := fmt.Sprintf("Failed to load configurations\n%v", err)
		panic(msg)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", controller.Index)

	r.Mount("/api", router())
	r.Mount("/admin", adminRouter())

	flag.Parse()
	if *gendoc {
		fmt.Println(docgen.JSONRoutesDoc(r))
	}

	http.ListenAndServe(":3333", r)
}
