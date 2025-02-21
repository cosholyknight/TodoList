package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Tasks interface {
		Create(ctx context.Context) error
	}

	Lists interface {
		Create(ctx context.Context) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Tasks: &TasksStore{db},
		Lists: &ListsStore{db},
	}
}
