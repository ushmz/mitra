package handler

import (
	"mitra/config"
	"net/http"

	"github.com/go-chi/render"
)

type Handler struct {
	Log    *LogHandler
	Search *SearchHandler
	User   *UserHandler
}

func NewHandler() *Handler {
	return &Handler{
		Log:    NewLogHandler(),
		Search: NewSearchHandler(),
		User:   NewUserHandler(),
	}
}

// Index return greeting response
func Index(w http.ResponseWriter, r *http.Request) {
	c := config.GetConfig()
	v := c.GetString("version")
	if v == "" {
		v = "beta"
	}
	greet := "" +
		`   __  ____ __           ` + "\n" +
		`  /  |/  (_) /________ _ ` + "\n" +
		` / /|_/ / / __/ __/ _ \/ ` + "\n" +
		`/_/  /_/_/\__/_/  \_,_/  ` + c.GetString("version") + "\n" +
		`Simple backend API server for the user-study.` + "\n" +
		`=================================================` + "\n"
	render.PlainText(w, r, greet)
}
