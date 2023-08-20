package user

import (
	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	upb "github.com/wizlif/tempmee_assignment/pb/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func dbUserToPbUser(user db.User) *upb.User {
	return &upb.User{
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
