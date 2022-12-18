-- name: CreateUser :exec
INSERT INTO users (
    username,
    password,
    email,
    sex,
    data_of_birth,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
);

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: LoginUser :one
SELECT * FROM users
WHERE username = $1 AND password = $2;

-- name: UpdateUser :exec
UPDATE users
SET username = $2,
    email = $3,
    updated_at = $4
WHERE username = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;