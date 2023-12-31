package user

import (
	"context"

	"github.com/mattn/go-sqlite3"
	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	upb "github.com/wizlif/tempmee_assignment/pb/user/v1"
	"github.com/wizlif/tempmee_assignment/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *UserServer) CreateUser(ctx context.Context, req *upb.CreateUserRequest) (*upb.CreateUserResponse, error) {
	if err := util.ValidateRequest(req); err != nil {
		return nil, err
	}

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account")
	}

	user, err := server.db.CreateUser(ctx, db.CreateUserParams{
		Email:    req.GetEmail(),
		Password: hashedPassword,
	})

	if err != nil {
		if sqlErr, ok := err.(sqlite3.Error); ok && sqlErr.Code == sqlite3.ErrConstraint {
			return nil, status.Errorf(codes.AlreadyExists, "account already exists")
		}
		
		return nil, status.Errorf(codes.Internal, "failed to create account")
	}

	return &upb.CreateUserResponse{
		User: dbUserToPbUser(user),
	}, nil
}
