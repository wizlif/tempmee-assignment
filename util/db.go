package util

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"

	"github.com/rs/zerolog/log"
)

func RunDBMigration(conn *sql.DB, dbSource string, dbMigration string) {
	driver, err := sqlite.WithInstance(conn, &sqlite.Config{})
	if err != nil {
		panic(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		dbMigration,
		dbSource,
		driver,
	)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}
