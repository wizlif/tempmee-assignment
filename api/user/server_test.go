package user

import (
	"testing"

	db "github.com/wizlif/tempmee_assignment/db/sqlc"
	"github.com/wizlif/tempmee_assignment/util"
)

func newTestUserServer(t *testing.T, store db.Store) *UserServer {
	config := util.Config{}

	server := &UserServer{
		config: config,
		db:     store,
	}

	return server
}
