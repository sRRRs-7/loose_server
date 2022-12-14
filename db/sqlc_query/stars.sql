-- name: CreateStar :exec
INSERT INTO stars (
    user_id, code_id
) VALUES (
    $1, $2
);

-- name: CountStar :one
SELECT COUNT(code_id) FROM stars
WHERE code_id = $1;

-- name: DeleteStar :exec
DELETE FROM stars
WHERE user_id = $1 AND code_id = $2;