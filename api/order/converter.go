package order

import (
	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	opb "github.com/wizlif/tempmee_assignment/pb/order/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func dbVOrderToPbOrder(order db.VOrder) *opb.Order {
	return &opb.Order{
		Id:   int32(order.ID),
		Book: order.Title,
		Status: order.Status,
		Amount: order.Price * float64(order.Total),
		Date: timestamppb.New(order.CreatedAt),
	}
}

func dbVOrdersToPbOrders(orders []db.VOrder) (pbOrders []*opb.Order) {
	for _, b := range orders {
		pbOrders = append(pbOrders, dbVOrderToPbOrder(b))
	}

	return
}
