package user

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	upb "github.com/wizlif/tempmee_assignment/pb/user/v1"
	"github.com/wizlif/tempmee_assignment/util"
)

type UserServer struct {
	upb.UnimplementedUserServiceServer
	config          util.Config
	db              db.Store
}

func RegisterUserGrpcServer(grpcMux *runtime.ServeMux, config util.Config,store db.Store) {
	server := &UserServer{
		config:          config,
		db:              store,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := upb.RegisterUserServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register user handler service")
	}
}
