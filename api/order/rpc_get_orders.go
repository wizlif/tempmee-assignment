package order

import (
	"context"

	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	opb "github.com/wizlif/tempmee_assignment/pb/order/v1"
	"github.com/wizlif/tempmee_assignment/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *OrderServer) GetOrders(ctx context.Context, req *opb.GetOrdersRequest) (*opb.GetOrdersResponse, error) {

	if err := util.ValidateRequest(req); err != nil {
		return nil, err
	}

	page := func() int64 {
		if req.Page == nil || req.GetPage() <= 0 {
			return 1
		}
		return req.GetPage()
	}()

	limit := func() int64 {
		if req.Limit == nil || req.GetLimit() <= 0 {
			return 10
		}
		return req.GetLimit()
	}()

	orders, err := server.db.ListUserOrders(ctx, db.ListUserOrdersParams{
		Offset: (page - 1) * limit,
		Limit:  limit,
		Email:  req.GetEmail(),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get orders")
	}

	return &opb.GetOrdersResponse{
		Orders: dbVOrdersToPbOrders(orders),
	}, nil
}
