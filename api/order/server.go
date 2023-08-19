package order

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	opb "github.com/wizlif/tempmee_assignment/pb/order/v1"
	"github.com/wizlif/tempmee_assignment/util"
)

type OrderServer struct {
	opb.UnimplementedOrderServiceServer
	config util.Config
	db     db.Store
}

func RegisterOrderGatewayServer(grpcMux *runtime.ServeMux, config util.Config, store db.Store) {
	server := &OrderServer{
		config: config,
		db:     store,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := opb.RegisterOrderServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register order handler service")
	}
}
