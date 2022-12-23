package domain

type Task struct {
	ID          int64  `db:"id" json:"id"`
	Topic       string `db:"topic" json:"topic"`
	Title       string `db:"title" json:"title"`
	Query       string `db:"query" json:"query"`
	Description string `db:"description" json:"description"`
}

type TaskSimple struct {
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
}

type CreateTaskParameters struct {
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
}

type AssignedTask struct {
	TaskID    int    `db:"task_id" json:"task_id"`
	Condition string `db:"condition" json:"condition"`
}

type TaskTopic struct {
	TaskID int    `db:"id" json:"id"`
	Topic  string `db:"topic" json:"topic"`
}

type TaskTopicUsed struct {
	Task1 bool `json:"task1"`
	Task2 bool `json:"task2"`
}

type Answer struct {
	UserID    int    `json:"user_id" db:"user_id"`
	TaskID    int    `json:"task_id" db:"task_id"`
	Condition string `json:"condition" db:"condition"`
	Answer    string `json:"answer" db:"answer"`
	Reason    string `json:"reason" db:"reason"`
}
