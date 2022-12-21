-- name: CreateCollection :exec
INSERT INTO collection (
    user_id, code_id
) VALUES (
    $1, $2
);

-- name: GetCollection :one
SELECT c.* FROM collection AS col
INNER JOIN users AS u ON col.user_id = u.id
INNER JOIN codes AS c ON col.code_id = c.id
WHERE c.id = $1;

-- name: GetAllCollections :many
SELECT DISTINCT c.id, col.id, c.* FROM collection AS col
INNER JOIN users AS u ON col.user_id = u.id
INNER JOIN codes AS c ON col.code_id = c.id
WHERE col.user_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetAllCollectionsBySearch :many
SELECT DISTINCT c.id, col.id, c.* FROM collection AS col
INNER JOIN users AS u ON col.user_id = u.id
INNER JOIN codes AS c ON col.code_id = c.id
WHERE col.user_id = $1 AND
    (
    c.username LIKE $2 OR
    c.code LIKE $3 OR
    c.description LIKE $4 OR
    $5 = ANY(tags)
    )
ORDER BY created_at DESC
LIMIT $6
OFFSET $7;

-- name: DeleteCollection :exec
DELETE FROM collection
WHERE id = $1;