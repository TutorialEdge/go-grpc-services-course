package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Migrate - handles migrations on our database
func (s *Store) Migrate() error {
	// we have a database connection setup within our store
	// we can reference the underlying sql.DB pointer using s.db.DB
	driver, err := postgres.WithInstance(s.db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	// here we want to run our migrations located within the
	// migrations directory. These migrations will create our tables
	// if they dont exist and perform actions such as adding new columns.
	// These will run in order and basically act as a list of steps that
	// eventually setup your database the way you need for production.
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	m.Steps(2)
	return nil
}
