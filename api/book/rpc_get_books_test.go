package book

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	mockdb "github.com/wizlif/tempmee_assignment/db/mock"
	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	bpb "github.com/wizlif/tempmee_assignment/pb/book/v1"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetBooksAPI(t *testing.T) {
	books := []db.Book{
		{
			ID:        1,
			Title:     "Living in the Light: A guide to personal transformation",
			Author:    "David Sargent",
			Price:     4.99,
			PageCount: 258,
		},
	}

	var page int64 = 1
	var limit int64 = 10

	testCases := []struct {
		name          string
		req           *bpb.GetBooksRequest
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res *bpb.GetBooksResponse, err error)
	}{
		{
			name: "OK",
			req:  &bpb.GetBooksRequest{},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListBooks(gomock.Any(), gomock.Any()).
					Times(1).
					Return(books, nil)
			},
			checkResponse: func(t *testing.T, res *bpb.GetBooksResponse, err error) {
				verifyBooksAreEqual(t, res, err, books)
			},
		}, {
			name: "OKWithPageAndLimit",
			req: &bpb.GetBooksRequest{
				Page:  &page,
				Limit: &limit,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListBooks(gomock.Any(), gomock.Any()).
					Times(1).
					Return(books, nil)
			},
			checkResponse: func(t *testing.T, res *bpb.GetBooksResponse, err error) {
				verifyBooksAreEqual(t, res, err, books)
			},
		},
		{
			name: "InternalError",
			req:  &bpb.GetBooksRequest{},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListBooks(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Book{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *bpb.GetBooksResponse, err error) {
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
			server := newTestBookServer(t, store)

			res, err := server.GetBooks(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}

// Verify db.Book[] == bpb.Book[]
func verifyBooksAreEqual(t *testing.T, res *bpb.GetBooksResponse, err error, books []db.Book) {
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, len(res.Books), len(books))

	for i, b := range res.Books {
		dbBook := dbBookToPbBook(books[i])
		require.Equal(t, b.Id, dbBook.Id)
		require.Equal(t, b.Title, dbBook.Title)
		require.Equal(t, b.Author, dbBook.Author)
		require.Equal(t, b.PageCount, dbBook.PageCount)
		require.Equal(t, b.Price, dbBook.Price)
	}
}
