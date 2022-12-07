package store

import (
	firebase "firebase.google.com/go"
	"github.com/jmoiron/sqlx"
)

// SQLStore has sql repositories
type Store struct {
	Auth   AuthenticationStore
	Log    LogStore
	Search SearchStore
	Task   TaskStore
	User   UserStore
}

// New return new SQLStore
func New(db *sqlx.DB, app *firebase.App) *Store {
	store := &Store{}
	store.Auth = NewAuthenticationStore(app)
	store.Log = NewLogStore(db)
	store.Search = NewSearchStore(db)
	store.Task = NewTaskStore(db)
	store.User = NewUserStore(db)
	return store
}
