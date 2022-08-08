package store

import "github.com/jmoiron/sqlx"

// SQLStore has sql repositories
type SQLStore struct {
	Log    LogStore
	Search SearchStore
	Task   TaskStore
	User   UserStore
}

// New return new SQLStore
func New(db *sqlx.DB) *SQLStore {
	store := &SQLStore{}
	store.Log = NewLogStore(db)
	store.Search = NewSearchStore(db)
	store.Task = NewTaskStore(db)
	store.User = NewUserStore(db)
	return store
}
