
-- name: ListProducts :many

SELECT * FROM products;

-- name: ListProductById :one
SELECT * FROM products WHERE id=$1;

-- name: CreateOrder :one
INSERT INTO orders ("userId", "total", "status", "address")
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items ("orderId", "productId", "quantity", "price")
VALUES ($1, $2, $3, $4)
RETURNING *;