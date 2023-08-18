package user

import (
	"context"

	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	upb "github.com/wizlif/tempmee_assignment/pb/user/v1"
	"github.com/wizlif/tempmee_assignment/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bufbuild/protovalidate-go"
)

func (server *UserServer) CreateUser(ctx context.Context, req *upb.CreateUserRequest) (*upb.CreateUserResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to initialize validator")
	}

	if err = v.Validate(req); err != nil {
		return nil, util.InvalidArgumentError("validation failed", util.ValidationErrorToFieldViolation(err))
	}

	_, err = server.db.CreateUser(ctx, db.CreateUserParams{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account")
	}

	return &upb.CreateUserResponse{}, nil
}
