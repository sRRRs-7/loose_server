// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: collection.sql

package db

import (
	"context"
	"time"
)

const createCollection = `-- name: CreateCollection :exec
INSERT INTO collection (
    user_id, code_id
) VALUES (
    $1, $2
)
`

type CreateCollectionParams struct {
	UserID int64 `json:"user_id"`
	CodeID int64 `json:"code_id"`
}

func (q *Queries) CreateCollection(ctx context.Context, arg CreateCollectionParams) error {
	_, err := q.db.Exec(ctx, createCollection, arg.UserID, arg.CodeID)
	return err
}

const deleteCollection = `-- name: DeleteCollection :exec
DELETE FROM collection
WHERE id = $1
`

func (q *Queries) DeleteCollection(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCollection, id)
	return err
}

const getAllCollections = `-- name: GetAllCollections :many
SELECT DISTINCT c.id, col.id, c.id, c.username, c.code, c.img, c.description, c.performance, c.star, c.tags, c.created_at, c.updated_at, c.access FROM collection AS col
INNER JOIN users AS u ON col.user_id = u.id
INNER JOIN codes AS c ON col.code_id = c.id
WHERE col.user_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type GetAllCollectionsParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetAllCollectionsRow struct {
	ID          int64     `json:"id"`
	ID_2        int64     `json:"id_2"`
	ID_3        int64     `json:"id_3"`
	Username    string    `json:"username"`
	Code        string    `json:"code"`
	Img         []byte    `json:"img"`
	Description string    `json:"description"`
	Performance string    `json:"performance"`
	Star        []int64   `json:"star"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Access      int64     `json:"access"`
}

func (q *Queries) GetAllCollections(ctx context.Context, arg GetAllCollectionsParams) ([]*GetAllCollectionsRow, error) {
	rows, err := q.db.Query(ctx, getAllCollections, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetAllCollectionsRow{}
	for rows.Next() {
		var i GetAllCollectionsRow
		if err := rows.Scan(
			&i.ID,
			&i.ID_2,
			&i.ID_3,
			&i.Username,
			&i.Code,
			&i.Img,
			&i.Description,
			&i.Performance,
			&i.Star,
			&i.Tags,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Access,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllCollectionsBySearch = `-- name: GetAllCollectionsBySearch :many
SELECT DISTINCT c.id, col.id, c.id, c.username, c.code, c.img, c.description, c.performance, c.star, c.tags, c.created_at, c.updated_at, c.access FROM collection AS col
INNER JOIN users AS u ON col.user_id = u.id
INNER JOIN codes AS c ON col.code_id = c.id
WHERE col.user_id = $1 AND
    (c.username LIKE $2 OR c.code LIKE $3 OR c.description LIKE $4)
ORDER BY created_at DESC
LIMIT $5
OFFSET $6
`

type GetAllCollectionsBySearchParams struct {
	UserID      int64  `json:"user_id"`
	Username    string `json:"username"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Limit       int32  `json:"limit"`
	Offset      int32  `json:"offset"`
}

type GetAllCollectionsBySearchRow struct {
	ID          int64     `json:"id"`
	ID_2        int64     `json:"id_2"`
	ID_3        int64     `json:"id_3"`
	Username    string    `json:"username"`
	Code        string    `json:"code"`
	Img         []byte    `json:"img"`
	Description string    `json:"description"`
	Performance string    `json:"performance"`
	Star        []int64   `json:"star"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Access      int64     `json:"access"`
}

func (q *Queries) GetAllCollectionsBySearch(ctx context.Context, arg GetAllCollectionsBySearchParams) ([]*GetAllCollectionsBySearchRow, error) {
	rows, err := q.db.Query(ctx, getAllCollectionsBySearch,
		arg.UserID,
		arg.Username,
		arg.Code,
		arg.Description,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetAllCollectionsBySearchRow{}
	for rows.Next() {
		var i GetAllCollectionsBySearchRow
		if err := rows.Scan(
			&i.ID,
			&i.ID_2,
			&i.ID_3,
			&i.Username,
			&i.Code,
			&i.Img,
			&i.Description,
			&i.Performance,
			&i.Star,
			&i.Tags,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Access,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCollection = `-- name: GetCollection :one
SELECT c.id, c.username, c.code, c.img, c.description, c.performance, c.star, c.tags, c.created_at, c.updated_at, c.access FROM collection AS col
INNER JOIN users AS u ON col.user_id = u.id
INNER JOIN codes AS c ON col.code_id = c.id
WHERE c.id = $1
`

func (q *Queries) GetCollection(ctx context.Context, id int64) (*Codes, error) {
	row := q.db.QueryRow(ctx, getCollection, id)
	var i Codes
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Code,
		&i.Img,
		&i.Description,
		&i.Performance,
		&i.Star,
		&i.Tags,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Access,
	)
	return &i, err
}
