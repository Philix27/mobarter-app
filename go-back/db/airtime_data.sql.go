// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: airtime_data.sql

package db

import (
	"context"
)

const listAirtime = `-- name: ListAirtime :many
SELECT id, network, country, created_at, updated_at FROM airtime
ORDER BY network
`

// ! #SECTION airtime
func (q *Queries) ListAirtime(ctx context.Context) ([]Airtime, error) {
	rows, err := q.db.Query(ctx, listAirtime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Airtime
	for rows.Next() {
		var i Airtime
		if err := rows.Scan(
			&i.ID,
			&i.Network,
			&i.Country,
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

const listDataPlans = `-- name: ListDataPlans :many
SELECT id, network, country, amount, duration, plan, is_supported, created_at, updated_at FROM data_plans
ORDER BY network
`

// ! #SECTION data_plans
func (q *Queries) ListDataPlans(ctx context.Context) ([]DataPlan, error) {
	rows, err := q.db.Query(ctx, listDataPlans)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DataPlan
	for rows.Next() {
		var i DataPlan
		if err := rows.Scan(
			&i.ID,
			&i.Network,
			&i.Country,
			&i.Amount,
			&i.Duration,
			&i.Plan,
			&i.IsSupported,
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
