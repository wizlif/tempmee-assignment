// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
)

type Querier interface {
	CreateBook(ctx context.Context, arg CreateBookParams) (int64, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (int64, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetBookById(ctx context.Context, id int64) (Book, error)
	GetOrderById(ctx context.Context, id int64) (VOrder, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error)
	ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error)
	ListUserOrders(ctx context.Context, arg ListUserOrdersParams) ([]VOrder, error)
	UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) error
}

var _ Querier = (*Queries)(nil)
