package order

import (
	"testing"

	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	"github.com/wizlif/tempmee_assignment/util"
)

func newTestOrderServer(t *testing.T, store db.Store) *OrderServer {
	config := util.Config{}

	server := &OrderServer{
		config: config,
		db:     store,
	}

	return server
}
