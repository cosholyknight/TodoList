package store

import (
	"context"
	"database/sql"
)

type Task struct {
	ID        int64  `json: "id"`
	Title     string `json: "title"`
	IsDone    bool   `json: "is_done"`
	ListID    int64  `json: "list_id"`
	CreatedAt string `json: "created_at"`
	UpdatedAt string `json: "updated_at"`
}
type TasksStore struct {
	db *sql.DB
}

func (s *TasksStore) Create(ctx context.Context, task *Task) error {
	query := `
		INSERT INTO tasks (title, is_done, list_id)
		VALUES ($1, $2, $3) RETURNING id, created_at, updated_at
    `
	err := s.db.QueryRowContext(
		ctx,
		query,
		task.Title,
		task.IsDone,
		task.ListID,
	).Scan(
		&task.ID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *TasksStore) GetTasksByListID(ctx context.Context, listID int64) ([]*Task, error) {
	query := `
		SELECT id, title, is_done, list_id, created_at, updated_at 
		FROM tasks 
		WHERE list_id = $1
	`
	rows, err := s.db.QueryContext(ctx, query, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		task := &Task{}
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.IsDone,
			&task.ListID,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TasksStore) Update(ctx context.Context, taskID int64, isDone bool) error {
	query := `
		UPDATE tasks 
		SET is_done = $1, updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
	`
	_, err := s.db.ExecContext(ctx, query, isDone, taskID)
	return err
}

func (s *TasksStore) Delete(ctx context.Context, taskID int64) error {
	query := `
		DELETE
		FROM tasks
		WHERE id = $1
`
	_, err := s.db.ExecContext(ctx, query, taskID)
	return err
}
