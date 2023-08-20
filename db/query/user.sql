-- name: CreateUser :one
INSERT INTO users (email,password) VALUES (?, ?) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?
