package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) MigrateDB() error {
	fmt.Println("migration of database started")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("could not create the postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		return fmt.Errorf("could not run migrations %w: ", err)
	}

	fmt.Println("Database migrated successfully!")
	return nil
}
