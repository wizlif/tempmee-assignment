package book

import (
	"testing"

	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	"github.com/wizlif/tempmee_assignment/util"
)

func newTestBookServer(t *testing.T, store db.Store) *BookServer {
	config := util.Config{}

	server := &BookServer{
		config: config,
		db:     store,
	}

	return server
}
