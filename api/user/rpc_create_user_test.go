package user

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	mockdb "github.com/wizlif/tempmee_assignment/db/mock"
	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	upb "github.com/wizlif/tempmee_assignment/pb/user/v1"
	"github.com/wizlif/tempmee_assignment/util"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/mattn/go-sqlite3"
)

func TestCreateUserAPI(t *testing.T) {
	email := "test@example.com"
	password := "23ABab#ee"
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user := db.User{
		ID:       1,
		Email:    email,
		Password: hashedPassword,
	}

	testCases := []struct {
		name          string
		req           *upb.CreateUserRequest
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res *upb.CreateUserResponse, err error)
	}{
		{
			name: "OK",
			req: &upb.CreateUserRequest{
				Password: password,
				Email:    user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, res *upb.CreateUserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				createdUser := res.GetUser()
				require.Equal(t, user.Email, createdUser.Email)
			},
		},
		{
			name: "InternalError",
			req: &upb.CreateUserRequest{
				Password: password,
				Email:    user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *upb.CreateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "DuplicateUsername",
			req: &upb.CreateUserRequest{
				Password: password,
				Email:    user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sqlite3.Error{
						Code: sqlite3.ErrConstraint,
					})
			},
			checkResponse: func(t *testing.T, res *upb.CreateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.AlreadyExists, st.Code())
			},
		},
		{
			name: "InvalidEmail",
			req: &upb.CreateUserRequest{
				Password: password,
				Email:    "invalid-email",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *upb.CreateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
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
			server := newTestUserServer(t, store)

			res, err := server.CreateUser(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
