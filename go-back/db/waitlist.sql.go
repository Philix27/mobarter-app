// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: waitlist.sql

package db

import (
	"context"
)

const createWaitlist = `-- name: CreateWaitlist :one
INSERT INTO waitlist (
  email
) VALUES (
  $1
)
RETURNING id, email, created_at, updated_at
`

func (q *Queries) CreateWaitlist(ctx context.Context, email string) (Waitlist, error) {
	row := q.db.QueryRow(ctx, createWaitlist, email)
	var i Waitlist
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listWaitlist = `-- name: ListWaitlist :many
SELECT id, email, created_at, updated_at FROM waitlist
ORDER BY email
`

// ! #SECTION waitlist
func (q *Queries) ListWaitlist(ctx context.Context) ([]Waitlist, error) {
	rows, err := q.db.Query(ctx, listWaitlist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Waitlist
	for rows.Next() {
		var i Waitlist
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
