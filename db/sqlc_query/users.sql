-- name: CreateUser :one
INSERT INTO users (
    user_id,
    password,
    username,
    email,
    sex,
    data_of_birth,
    destination_family_name,
    destination_first_name,
    postcode,
    prefecture_code,
    city,
    street,
    building,
    phone,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
) RETURNING *;

-- name: GetUser :one
SELECT id FROM users
WHERE user_id = $1;

-- name: LoginUser :one
SELECT * FROM users
WHERE user_id = $1 AND password = $2;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :exec
UPDATE users
SET username = $2,
    email = $3,
    updated_at = $4
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;