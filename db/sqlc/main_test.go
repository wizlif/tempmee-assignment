package db

import (
	"database/sql"
	"os"
	"testing"

	"github.com/rs/zerolog/log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wizlif/tempmee_assignment/util"
)

var testStore Store

func TestMain(m *testing.M) {
	dbName := ":memory:"

	conn, err := sql.Open("sqlite3", dbName)
	if err != nil && conn.Ping() != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	util.RunDBMigration(conn, dbName, "file://../migration")

	testStore = NewStore(conn)
	os.Exit(m.Run())
}
