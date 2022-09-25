package domain

import (
	"time"
)

type CompletionCode struct {
	UserID    int       `db:"user_id"`
	Code      int       `db:"completion_code"`
	CreatedAT time.Time `db:"created_at" goqu:"defaultifempty"`
	UpdateAT  time.Time `db:"update_at" goqu:"defaultifempty"`
}
