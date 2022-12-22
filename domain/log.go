package domain

// DwellTimeLog : Struct for SERP dwell time log
type DwellTimeLog struct {
	// UserID : The ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// Condition : User's condition ID that means group and task category.
	Condition string `db:"condition" json:"condition"`

	// DwellTime : How many times the user spend in SERP
	DwellTime int `db:"time_on_page" json:"dwell_time"`
}

// ClickLog : Struct for SERP click log
type ClickLog struct {
	// UserID : The ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// Condition : User's condition ID that means group and task category.
	Condition string `db:"condition" json:"condition"`

	// Time : User's page viewing time.
	Time int `db:"time_on_page" json:"time"`

	// Rank : Search result rank that user clicked.
	Rank int `db:"ranked_at" json:"rank"`

	// IsVisible : Risk is visible or not.
	IsVisible bool `db:"is_visible" json:"visible"`

	// IsFirstClick : The click event is the first click or not
	IsFirstClick bool `db:"is_first" json:"is_first"`

	Event string `db:"event" json:"-"`
}
