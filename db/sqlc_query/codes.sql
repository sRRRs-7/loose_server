-- name: CreateCode :exec
INSERT INTO codes (
    username, code, img, description, performance, star, tags, created_at, updated_at, access
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
);

-- name: GetCode :one
SELECT * FROM codes
WHERE id = $1;

-- name: GetAllCodes :many
SELECT * FROM codes
ORDER BY created_at ASC
LIMIT $1
OFFSET $2;

-- name: GetAllCodesByTag :many
SELECT * FROM codes
WHERE
    $1 = ANY(tags) OR
    $2 = ANY(tags) AND
    $3 = ANY(tags) AND
    $4 = ANY(tags) AND
    $5 = ANY(tags) AND
    $6 = ANY(tags) AND
    $7 = ANY(tags) AND
    $8 = ANY(tags) AND
    $9 = ANY(tags) AND
    $10 = ANY(tags)
ORDER BY created_at ASC
LIMIT $11
OFFSET $12;

-- name: GetAllCodesByKeyword :many
SELECT * FROM codes
WHERE
    username LIKE $1 OR
    code LIKE $2 OR
    description LIKE $3
ORDER BY created_at ASC
LIMIT $4
OFFSET $5;

-- name: GetAllCodesSortedStar :many
SELECT * FROM codes
ORDER BY array_length(star, 1) ASC
LIMIT $1
OFFSET $2;

-- name: GetAllCodesSortedAccess :many
SELECT * FROM codes
ORDER BY access DESC
LIMIT $1
OFFSET $2;

-- name: GetAllOwnCodes :many
SELECT * FROM codes
WHERE username = $1
LIMIT $2
OFFSET $3;

-- name: UpdateCode :exec
UPDATE codes
SET code = $2,
    img = $3,
    description = $4,
    performance = $5,
    tags = $6,
    updated_at = $7
WHERE id = $1;

-- name: UpdateAccess :exec
UPDATE codes
SET access = $2
WHERE id = $1;

-- name: UpdateStar :exec
UPDATE codes
SET star = $2
WHERE id = $1;

-- name: DeleteCode :exec
DELETE FROM codes
WHERE id = $1;