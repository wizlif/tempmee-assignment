-- name: CreateOrder :one
INSERT INTO orders (price,book_id,user_id,total,status) VALUES (?, ?, ?, ?, ?) RETURNING id;

-- name: ListOrders :many
SELECT * FROM orders ORDER BY id LIMIT ? OFFSET ?;

-- name: ListUserOrders :many
SELECT * FROM v_orders WHERE email = ? ORDER BY id LIMIT ? OFFSET ?;

-- name: GetOrderById :one
SELECT * FROM v_orders WHERE id = ?;

-- name: UpdateOrderStatus :exec
UPDATE orders SET status = ? WHERE id = ?;
