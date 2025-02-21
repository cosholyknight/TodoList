package store

import (
	"context"
	"database/sql"
)

type ListsStore struct {
	db *sql.DB
}

func (s *ListsStore) Create(ctx context.Context) error {
	return nil
}
