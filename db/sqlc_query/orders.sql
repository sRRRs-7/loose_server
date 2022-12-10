-- name: CreateOrder :one
INSERT INTO orders (
    user_id, product_id, quantity, postage, price, status, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE user_id = $1 and product_id = $2;

-- name: ListOrders :many
SELECT * FROM orders
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateOrderQuantity :exec
UPDATE orders
SET quantity = $2
where id = $1
RETURNING *;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;