package main

import (
	"context"
	"fmt"

	"github.com/melinaco4/companies-manager/internal/db"
)

func Run() error {
	fmt.Println("Start up app")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Health Check: Pinged database successfully")

	if err := db.MigrateDB(); err != nil {
		return err
	}
	fmt.Println("Migration Check: Migrated database successfully")

	return nil
}
func main() {
	fmt.Println("Heyyyyyy!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
