-- name: CreateBook :one
INSERT INTO books (title,author,price,page_count) VALUES (?, ?, ?, ?) RETURNING id;

-- name: ListBooks :many
SELECT * FROM books ORDER BY id LIMIT ? OFFSET ?;
