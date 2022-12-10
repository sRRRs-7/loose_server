
-- name: CreateAdminUser :one
INSERT INTO adminuser(
    username, password, created_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetAdminUser :one
SELECT * FROM adminuser WHERE username = $1 AND password = $2;