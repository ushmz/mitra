package handler

import (
	"mitra/config"
	"mitra/store"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Log    *LogHandler
	Search *SearchHandler
	Task   *TaskHandler
	User   *UserHandler
}

func NewHandler(db *sqlx.DB, app *firebase.App) *Handler {
	s := store.New(db, app)
	return &Handler{
		Log:    NewLogHandler(),
		Task:   NewTaskHandler(s),
		Search: NewSearchHandler(s),
		User:   NewUserHandler(s),
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
