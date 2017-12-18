package postgres

import (
	"database/sql"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"github.com/mattes/migrate/source/go-bindata"

	"checker/storage/postgres/migrations"
)

func initMigrations(dbName string, db *sql.DB) error {
	// wrap assets into Resource
	resource := bindata.Resource(
		migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		},
	)

	sourceDriver, err := bindata.WithInstance(resource)
	if err != nil {
		return err
	}

	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	migrates, err := migrate.NewWithInstance("go-bindata", sourceDriver, dbName, dbDriver)
	if err != nil {
		return err
	}

	if err := migrates.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
