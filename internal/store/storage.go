package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Tasks interface {
		Create(ctx context.Context, task *Task) error
	}

	Lists interface {
		Create(ctx context.Context, list *List) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Tasks: &TasksStore{db},
		Lists: &ListsStore{db},
	}
}
