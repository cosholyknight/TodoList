package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Tasks interface {
		Create(ctx context.Context, task *Task) error
		GetTasksByListID(ctx context.Context, listID int64) ([]*Task, error)
		Update(ctx context.Context, taskID int64, isDone bool) error
		Delete(ctx context.Context, taskID int64) error
	}

	Lists interface {
		Create(ctx context.Context, list *List) error
		Delete(ctx context.Context, listID int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Tasks: &TasksStore{db},
		Lists: &ListsStore{db},
	}
}
