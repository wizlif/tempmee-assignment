package order

import (
	"context"
	"database/sql"

	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	opb "github.com/wizlif/tempmee_assignment/pb/order/v1"
	"github.com/wizlif/tempmee_assignment/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *OrderServer) CreateOrder(ctx context.Context, req *opb.CreateOrderRequest) (*opb.CreateOrderResponse, error) {

	if err := util.ValidateRequest(req); err != nil {
		return nil, err
	}

	user, err := server.db.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.FailedPrecondition, "no such user")
		}
		return nil, status.Errorf(codes.Internal, "failed to create order")
	}

	book, err := server.db.GetBookById(ctx, req.GetBookId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.FailedPrecondition, "no such book")
		}
		return nil, status.Errorf(codes.Internal, "failed to create order")
	}

	orderId, err := server.db.CreateOrder(ctx, db.CreateOrderParams{
		Price:  book.Price,
		BookID: book.ID,
		UserID: user.ID,
		Total:  req.GetTotal(),
		Status: "PENDING",
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create order")
	}

	order, err := server.db.GetOrderById(ctx, orderId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve order")
	}

	return &opb.CreateOrderResponse{
		Order: dbVOrderToPbOrder(order),
	}, nil
}
