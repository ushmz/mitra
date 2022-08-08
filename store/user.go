package store

import (
	"context"
	"mitra/domain/model"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type UserStore interface {
	CreateUser(context.Context, model.User) error
	IssueCompletionCode(context.Context, model.CompletionCode) error
}

type UserStoreImple struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) UserStore {
	return &UserStoreImple{db: db}
}

func (s *UserStoreImple) CreateUser(ctx context.Context, u model.User) error {
	if s == nil {
		return ErrNilReceiver
	}

	q, a, err := dialect.Insert("users").Rows(u).ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a...); err != nil {
		return ErrDatabaseExecutionFailere
	}

	return nil
}

func (s *UserStoreImple) UpdateUserUID(ctx context.Context, userID int, uid string) error {
	if s == nil {
		return ErrNilReceiver
	}

	q, a, err := dialect.Update("users").Set(goqu.Record{"uid": uid}).Where(goqu.C("user_id").Eq(userID)).ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a...); err != nil {
		return ErrDatabaseExecutionFailere
	}

	return nil
}

func (s *UserStoreImple) IssueCompletionCode(ctx context.Context, c model.CompletionCode) error {
	if s == nil {
		return ErrNilReceiver
	}

	q, a, err := dialect.Insert("completion_codes").Rows(c).ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a...); err != nil {
		return ErrDatabaseExecutionFailere
	}
	return nil
}
