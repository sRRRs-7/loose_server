-- name: CreateCategory :one
INSERT INTO categories (
    category
) VALUES (
    $1
) RETURNING *;

-- name: GetCategories :one
SELECT * FROM categories
WHERE category = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY id
LIMIT $1
OFFSET $2;
