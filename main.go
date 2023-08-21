package main

import (
	"net"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/wizlif/tempmee_assignment/api"
	"github.com/wizlif/tempmee_assignment/api/book"
	"github.com/wizlif/tempmee_assignment/api/order"
	"github.com/wizlif/tempmee_assignment/api/user"
	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	"github.com/wizlif/tempmee_assignment/util"

	"database/sql"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/wizlif/tempmee_assignment/doc/statik"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open("sqlite3", config.DBName)
	if err != nil && conn.Ping() != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	util.RunDBMigration(conn, config.DBName, "file://db/migration")

	store := db.NewStore(conn)

	runGatewayServer(config, store)
}

func runGatewayServer(config util.Config, store db.Store) {

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	user.RegisterUserGatewayServer(grpcMux, config, store)
	book.RegisterBookGatewayServer(grpcMux, config, store)
	order.RegisterOrderGatewayServer(grpcMux, config, store)

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik fs")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress())
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("started HTTP gateway server at %s", listener.Addr().String())
	handler := api.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}
}
