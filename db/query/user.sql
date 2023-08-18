-- name: CreateUser :one
INSERT INTO users (email,password) VALUES (?, ?) RETURNING id;
