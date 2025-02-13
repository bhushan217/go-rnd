// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: role.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRole = `-- name: CreateRole :one
INSERT INTO roles (
  title
) VALUES (
  $1
)
RETURNING id, title, created_at, created_by, updated_at, updated_by, version, latest
`

func (q *Queries) CreateRole(ctx context.Context, id string) (Role, error) {
	row := q.db.QueryRow(ctx, createRole, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Version,
		&i.Latest,
	)
	return i, err
}

const deleteRole = `-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1 and version = $2
`

type DeleteRoleParams struct {
	ID      int64       `db:"id" json:"id"`
	Version pgtype.Int8 `db:"version" json:"version"`
}

func (q *Queries) DeleteRole(ctx context.Context, arg DeleteRoleParams) error {
	_, err := q.db.Exec(ctx, deleteRole, arg.ID, arg.Version)
	return err
}

const getRole = `-- name: GetRole :one
SELECT id, title, created_at, created_by, updated_at, updated_by, version, latest FROM roles
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRole(ctx context.Context, id int64) (Role, error) {
	row := q.db.QueryRow(ctx, getRole, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Version,
		&i.Latest,
	)
	return i, err
}

const listRole = `-- name: ListRole :many
SELECT id, title, created_at, created_by, updated_at, updated_by, version, latest FROM roles
ORDER BY title
`

func (q *Queries) ListRole(ctx context.Context) ([]Role, error) {
	rows, err := q.db.Query(ctx, listRole)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Role{}
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.Version,
			&i.Latest,
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

const updateRole = `-- name: UpdateRole :exec
UPDATE roles
  set title = $1
WHERE id = $2 and version = $3
`

type UpdateRoleParams struct {
	Title   string      `db:"title" json:"title"`
	ID      int64       `db:"id" json:"id"`
	Version pgtype.Int8 `db:"version" json:"version"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) error {
	_, err := q.db.Exec(ctx, updateRole, arg.Title, arg.ID, arg.Version)
	return err
}
