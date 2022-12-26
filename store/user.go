package store

import (
	"context"
	"database/sql"
	"mitra/domain"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type UserStore interface {
	CreateUser(context.Context, *domain.ImplicitUser) (*domain.UserSimple, error)
	GetCompletionCode(ctx context.Context, userID int) (int, error)
	SetCompletionCode(context.Context, *domain.CompletionCode) error
}

type UserStoreImple struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) UserStore {
	return &UserStoreImple{db: db}
}

func (s *UserStoreImple) CreateUser(ctx context.Context, u *domain.ImplicitUser) (*domain.UserSimple, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.
		Insert("users").
		Rows(u).
		Prepared(true).
		ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	if rs, err := s.db.ExecContext(ctx, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	} else {
		id, err := rs.LastInsertId()
		if err != nil {
			return nil, ErrDatabaseExecutionFailere
		}
		return &domain.UserSimple{ID: int(id), ExternalID: u.ExternalID}, nil
	}
}

func (s *UserStoreImple) UpdateUserUID(ctx context.Context, userID int, uid string) error {
	if s == nil {
		return ErrNilReceiver
	}

	q, a, err := dialect.
		Update("users").
		Set(goqu.Record{"uid": uid}).
		Where(goqu.C("user_id").Eq(userID)).
		Prepared(true).
		ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a...); err != nil {
		return ErrDatabaseExecutionFailere
	}

	return nil
}

func (s *UserStoreImple) GetCompletionCode(ctx context.Context, userID int) (int, error) {
	if s == nil {
		return 0, ErrNilReceiver
	}

	q, a, err := dialect.
		Select("completion_code").
		From("completion_codes").
		Where(goqu.C("user_id").Eq(userID)).
		Prepared(true).
		ToSQL()
	if err != nil {
		return 0, ErrQueryBuildFailure
	}

	var dest int
	if err := s.db.GetContext(ctx, &dest, q, a...); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, ErrDatabaseExecutionFailere
	}
	return dest, nil

}

func (s *UserStoreImple) SetCompletionCode(ctx context.Context, c *domain.CompletionCode) error {
	if s == nil {
		return ErrNilReceiver
	}

	q, a, err := dialect.
		Insert("completion_codes").
		Rows(goqu.Record{
			"user_id":         c.UserID,
			"completion_code": c.Code,
		}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a...); err != nil {
		return ErrDatabaseExecutionFailere
	}
	return nil
}
