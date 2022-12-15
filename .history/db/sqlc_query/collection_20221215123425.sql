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
SELECT c.*, col.id FROM collection AS col
INNER JOIN users AS u ON col.user_id = u.id
INNER JOIN codes AS c ON col.code_id = c.id
WHERE col.user_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetAllCollectionsBySearch :many
SELECT c.*, col.id FROM collection AS col
INNER JOIN users AS u ON col.user_id = u.id
INNER JOIN codes AS c ON col.code_id = c.id
WHERE col.user_id = $1 OR
        c.username LIKE $1 OR
    c.code LIKE $2 OR
    c.description LIKE $3
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: DeleteCollection :exec
DELETE FROM collection
WHERE id = $1;