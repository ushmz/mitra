package model

type Task struct {
	ID          int64  `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
}

type CreateTaskParameters struct {
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
}

type TaskSimple struct {
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
}
