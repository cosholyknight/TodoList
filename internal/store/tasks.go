package store

import (
	"context"
	"database/sql"
)

type TasksStore struct {
	db *sql.DB
}

func (s *TasksStore) Create(ctx context.Context) error {
	return nil
}
