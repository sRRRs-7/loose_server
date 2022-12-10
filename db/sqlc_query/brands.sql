-- name: CreateBrand :one
INSERT INTO brands (
    brand_name
) VALUES (
    $1
) RETURNING *;

-- name: Listbrands :many
SELECT * FROM brands
ORDER BY id;