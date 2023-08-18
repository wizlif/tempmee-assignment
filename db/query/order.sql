-- name: CreateOrder :one
INSERT INTO orders (price,book_id,user_id,status) VALUES (?, ?, ?, ?) RETURNING id;

-- name: ListOrders :many
SELECT * FROM orders ORDER BY id LIMIT ? OFFSET ?;

-- name: UpdateOrderStatus :exec
UPDATE orders SET status = ? WHERE id = ?;
