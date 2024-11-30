// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: country.sql

package db

import (
	"context"
)

const listCountry = `-- name: ListCountry :many
SELECT id, name, created_at, updated_at FROM country
ORDER BY name
`

// ! #SECTION country
func (q *Queries) ListCountry(ctx context.Context) ([]Country, error) {
	rows, err := q.db.Query(ctx, listCountry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Country
	for rows.Next() {
		var i Country
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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
