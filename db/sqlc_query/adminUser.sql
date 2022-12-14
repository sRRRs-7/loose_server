-- name: CreateAdminUser :exec
INSERT INTO adminuser(
    username, password, created_at
) VALUES (
    $1, $2, $3
);

-- name: GetAdminUser :one
SELECT * FROM adminuser
WHERE username = $1 AND password = $2;