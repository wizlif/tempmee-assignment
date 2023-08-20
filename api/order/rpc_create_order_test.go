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

func TestCreateOrderAPI(t *testing.T) {
	email := "test@example.com"
	var total int64 = 2

	book := db.Book{
		ID:        1,
		Title:     "Living in the Light: A guide to personal transformation",
		Author:    "David Sargent",
		Price:     4.99,
		PageCount: 258,
	}

	order := db.VOrder{
		ID:     1,
		Price:  39.99,
		Total:  total,
		Status: "PENDING",
		Title:  "Go Programming by Example",
		Email:  email,
	}

	user := db.User{
		ID:       1,
		Email:    email,
		Password: "abcd",
	}

	testCases := []struct {
		name          string
		req           *opb.CreateOrderRequest
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res *opb.CreateOrderResponse, err error)
	}{
		{
			name: "OK",
			req: &opb.CreateOrderRequest{
				Email:  email,
				BookId: book.ID,
				Total:  total,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), email).
					Times(1).
					Return(user, nil)

				store.EXPECT().
					GetBookById(gomock.Any(), book.ID).
					Times(1).
					Return(book, nil)

				arg := db.CreateOrderParams{
					Price:  book.Price,
					BookID: book.ID,
					UserID: user.ID,
					Total:  total,
					Status: "PENDING",
				}

				store.EXPECT().
					CreateOrder(gomock.Any(), arg).
					Times(1).
					Return(order.ID, nil)

				store.EXPECT().
					GetOrderById(gomock.Any(), order.ID).
					Times(1).
					Return(order, nil)

			},
			checkResponse: func(t *testing.T, res *opb.CreateOrderResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, int64(res.Order.Id), order.ID)
				// TODO: Add more checks
			},
		},
		{
			name: "ValidationError",
			req:  &opb.CreateOrderRequest{},
			buildStubs: func(store *mockdb.MockStore) {

			},
			checkResponse: func(t *testing.T, res *opb.CreateOrderResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "InternalErrorNoUser",
			req: &opb.CreateOrderRequest{
				Email:  email,
				BookId: book.ID,
				Total:  total,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *opb.CreateOrderResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "InternalErrorFailedToGetUser",
			req: &opb.CreateOrderRequest{
				Email:  email,
				BookId: book.ID,
				Total:  total,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *opb.CreateOrderResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "InternalErrorNoUser",
			req: &opb.CreateOrderRequest{
				Email:  email,
				BookId: book.ID,
				Total:  total,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, res *opb.CreateOrderResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.FailedPrecondition, st.Code())
			},
		},
		{
			name: "InternalErrorFailedToGetBook",
			req: &opb.CreateOrderRequest{
				Email:  email,
				BookId: book.ID,
				Total:  total,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)

				store.EXPECT().
					GetBookById(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Book{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *opb.CreateOrderResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "InternalErrorNoBook",
			req: &opb.CreateOrderRequest{
				Email:  email,
				BookId: book.ID,
				Total:  total,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)

				store.EXPECT().
					GetBookById(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Book{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, res *opb.CreateOrderResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.FailedPrecondition, st.Code())
			},
		},
		{
			name: "InternalErrorCreatingOrder",
			req: &opb.CreateOrderRequest{
				Email:  email,
				BookId: book.ID,
				Total:  total,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)

				store.EXPECT().
					GetBookById(gomock.Any(), gomock.Any()).
					Times(1).
					Return(book, nil)

				store.EXPECT().
					CreateOrder(gomock.Any(), gomock.Any()).
					Times(1).
					Return(order.ID, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *opb.CreateOrderResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "InternalErrorRetrievingOrder",
			req: &opb.CreateOrderRequest{
				Email:  email,
				BookId: book.ID,
				Total:  total,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)

				store.EXPECT().
					GetBookById(gomock.Any(), gomock.Any()).
					Times(1).
					Return(book, nil)

				store.EXPECT().
					CreateOrder(gomock.Any(), gomock.Any()).
					Times(1).
					Return(order.ID, nil)

				store.EXPECT().
					GetOrderById(gomock.Any(), gomock.Any()).
					Times(1).
					Return(order, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *opb.CreateOrderResponse, err error) {
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

			res, err := server.CreateOrder(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
