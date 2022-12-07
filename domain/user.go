package domain

type ImplicitUser struct {
	ExternalID  string `db:"external_id"`
	FirebaseUID string `db:"firebase_uid"`
}

type UserSimple struct {
	ID         int    `json:"id" db:"id"`
	ExternalID string `json:"external_id" db:"external_id"`
}

type User struct {
	ID         int    `json:"id" db:"id"`
	ExternalID string `json:"external_id" db:"external_id"`
	Token      string `json:"token" db:"-"`
}

type AssignedUser struct {
	ID         int    `json:"id" db:"id"`
	ExternalID string `json:"external_id" db:"external_id"`
	TaskID     int    `json:"task_id" db:"task_id"`
	Condition  string `json:"condition" db:"condition"`
}
