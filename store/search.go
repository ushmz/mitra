package store

import (
	"mitra/domain/model"

	"github.com/jmoiron/sqlx"
)

type SearchStore interface {
	GetSearchResults(offset, limit int) ([]model.SearchResult, error)
}

type SearchStoreImpl struct {
	db *sqlx.DB
}

func NewSearchStore(db *sqlx.DB) SearchStore {
	return &SearchStoreImpl{db: db}
}

func (s *SearchStoreImpl) GetSearchResults(offset, limit int) ([]model.SearchResult, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.
		Select("*").
		From("search_pages").
		Limit(uint(limit)).
		Offset(uint(offset * limit)).
		ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs := make([]model.SearchResult, limit)
	if err := s.db.Select(&rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}
