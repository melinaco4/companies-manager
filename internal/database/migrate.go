package database

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

//Migrating the database, reading from the files in the companies-manager/migration folder to create new table
// if it doesn't exist or drop the table.
//Irrelevant Comment:
// I also tried to create a MongoDB database with the gorm package and migrate it with the AutoMigrate function
// which was a lot easier but I had to change strategy as I had a hard time with the uuid id in the MongoDB
func (d *Database) MigrateDB() error {
	fmt.Println("migrating our database")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create the postgres driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("could not run up migrations: %w", err)
		}
	}

	fmt.Println("successfully migrated the database")
	return nil
}
