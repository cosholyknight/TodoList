package store

import (
	"context"
	"database/sql"
)

type List struct {
	ID        int64  `json: "id"`
	Title     string `json: "title"`
	CreatedAt string `json: "created_at"`
}
type ListsStore struct {
	db *sql.DB
}

func (s *ListsStore) Create(ctx context.Context, list *List) error {
	query := `
		INSERT INTO lists (title)
		VALUES ($1) RETURNING id, created_at
    `
	err := s.db.QueryRowContext(
		ctx,
		query,
		list.Title,
	).Scan(
		&list.ID,
		&list.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *ListsStore) Delete(ctx context.Context, listID int64) error {
	query := `
		DELETE
		FROM lists
		WHERE id = $1
`
	_, err := s.db.ExecContext(ctx, query, listID)
	if err != nil {
		return err
	}

	return nil
}
