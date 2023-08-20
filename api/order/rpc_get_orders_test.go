package order

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	mockdb "github.com/wizlif/tempmee_assignment/db/mock"
	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	opb "github.com/wizlif/tempmee_assignment/pb/order/v1"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetOrdersAPI(t *testing.T) {
	email := "test@example.com"
	orders := []db.VOrder{
		{
			ID:     1,
			Price:  39.99,
			Total:  2,
			Status: "PENDING",
			Title:  "Go Programming by Example",
			Email:  email,
		},
	}

	var page int64 = 1
	var limit int64 = 10

	testCases := []struct {
		name          string
		req           *opb.GetOrdersRequest
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res *opb.GetOrdersResponse, err error)
	}{
		{
			name: "OK",
			req: &opb.GetOrdersRequest{
				Email: email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListUserOrders(gomock.Any(), gomock.Any()).
					Times(1).
					Return(orders, nil)
			},
			checkResponse: func(t *testing.T, res *opb.GetOrdersResponse, err error) {
				verifyOrdersAreEqual(t, res, err, orders)
			},
		}, {
			name: "OKWithPageAndLimit",
			req: &opb.GetOrdersRequest{
				Page:  &page,
				Limit: &limit,
				Email: email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListUserOrders(gomock.Any(), gomock.Any()).
					Times(1).
					Return(orders, nil)
			},
			checkResponse: func(t *testing.T, res *opb.GetOrdersResponse, err error) {
				verifyOrdersAreEqual(t, res, err, orders)
			},
		},
		{
			name: "ValidationError",
			req: &opb.GetOrdersRequest{
			},
			buildStubs: func(store *mockdb.MockStore) {
				
			},
			checkResponse: func(t *testing.T, res *opb.GetOrdersResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "InternalError",
			req: &opb.GetOrdersRequest{
				Email: email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListUserOrders(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.VOrder{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *opb.GetOrdersResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			defer storeCtrl.Finish()
			store := mockdb.NewMockStore(storeCtrl)

			tc.buildStubs(store)
			server := newTestOrderServer(t, store)

			res, err := server.GetOrders(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}

// Verify db.Order[] == opb.Order[]
func verifyOrdersAreEqual(t *testing.T, res *opb.GetOrdersResponse, err error, orders []db.VOrder) {
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, len(res.Orders), len(orders))

	for i, b := range res.Orders {
		dbOrder := orders[i]
		require.Equal(t, int64(b.Id), dbOrder.ID)
		require.Equal(t, b.Book, dbOrder.Title)
		require.Equal(t, b.Amount, dbOrder.Price*float64(dbOrder.Total))
		require.Equal(t, b.Status, dbOrder.Status)
	}
}
