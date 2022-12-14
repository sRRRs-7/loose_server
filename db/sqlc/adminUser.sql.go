// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: adminUser.sql

package db

import (
	"context"
	"time"
)

const createAdminUser = `-- name: CreateAdminUser :exec
INSERT INTO adminuser(
    username, password, created_at
) VALUES (
    $1, $2, $3
)
`

type CreateAdminUserParams struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateAdminUser(ctx context.Context, arg CreateAdminUserParams) error {
	_, err := q.db.Exec(ctx, createAdminUser, arg.Username, arg.Password, arg.CreatedAt)
	return err
}

const getAdminUser = `-- name: GetAdminUser :one
SELECT id, username, password, created_at FROM adminuser
WHERE username = $1 AND password = $2
`

type GetAdminUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) GetAdminUser(ctx context.Context, arg GetAdminUserParams) (*Adminuser, error) {
	row := q.db.QueryRow(ctx, getAdminUser, arg.Username, arg.Password)
	var i Adminuser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return &i, err
}
