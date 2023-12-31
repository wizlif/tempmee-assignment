// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"database/sql"
	"time"
)

type Book struct {
	ID        int64        `json:"id"`
	Title     string       `json:"title"`
	Author    string       `json:"author"`
	Price     float64      `json:"price"`
	PageCount int64        `json:"page_count"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Order struct {
	ID        int64        `json:"id"`
	Price     float64      `json:"price"`
	BookID    int64        `json:"book_id"`
	UserID    int64        `json:"user_id"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	Total     int64        `json:"total"`
}

type User struct {
	ID        int64        `json:"id"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type VOrder struct {
	ID        int64     `json:"id"`
	Price     float64   `json:"price"`
	Total     int64     `json:"total"`
	Status    string    `json:"status"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
