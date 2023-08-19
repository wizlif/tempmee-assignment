package book

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	bpb "github.com/wizlif/tempmee_assignment/pb/book/v1"
	"github.com/wizlif/tempmee_assignment/util"
)

type BookServer struct {
	bpb.UnimplementedBookServiceServer
	config util.Config
	db     db.Store
}

func RegisterBookGatewayServer(grpcMux *runtime.ServeMux, config util.Config, store db.Store) {
	server := &BookServer{
		config: config,
		db:     store,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bpb.RegisterBookServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register book handler service")
	}
}
