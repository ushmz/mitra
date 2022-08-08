package store

import "github.com/jmoiron/sqlx"

type TaskStore interface{}

type TaskStoreImpl struct {
	db *sqlx.DB
}

func NewTaskStore(db *sqlx.DB) TaskStore {
	return &TaskStoreImpl{db: db}
}
