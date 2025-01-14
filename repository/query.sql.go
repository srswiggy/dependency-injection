// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package repository

import (
	"context"
)

const getUsers = `-- name: GetUsers :many
SELECT id, name FROM user_table
`

func (q *Queries) GetUsers(ctx context.Context) ([]UserTable, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserTable
	for rows.Next() {
		var i UserTable
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
