-- name: CreateMedia :one
INSERT INTO media (
    title, contents, img, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetMedia :one
SELECT * FROM media
WHERE id = $1 LIMIT 1;

-- name: ListMedia :many
SELECT * FROM media
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: UpdateMedia :exec
UPDATE media
SET title = $2,
    contents = $3,
    img = $4
WHERE id = $1
RETURNING *;

-- name: DeleteMedia :exec
DELETE FROM media
WHERE id = $1;