-- name: CreateProduct :exec
INSERT INTO products (
    product_name, description, img, unit_price, discount, stock, brand_id, category, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProduct :many
SELECT * FROM products
ORDER BY created_at ASC
LIMIT $1
OFFSET $2;

-- name: ListProductByKeyword :many
SELECT * FROM products
WHERE description LIKE $1
OR product_name LIKE $2
ORDER BY unit_price ASC
LIMIT $3
OFFSET $4;

-- name: ListProductByCategory :many
SELECT p.* FROM
products AS p
INNER JOIN
categories AS c
ON p.category = c.id
WHERE c.category = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: UpdateProductStock :one
UPDATE products
SET stock = $2,
    updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;