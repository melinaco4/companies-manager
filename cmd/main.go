package main

import (
	"context"
	"fmt"

	"github.com/melinaco4/companies-manager/internal/company"
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

	//cmpnyService := company.NewService(db)
	test, err := db.PostCompany(
		context.Background(),
		company.Company{
			ID:                "kj12sd23-2345-gfds-gv34-gr4egr24ff2d",
			Name:              "Test",
			Description:       "Just a description",
			AmountofEmployees: "1434",
			Registered:        true,
			Type:              "NonProfit",
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test)
	fmt.Println(db.GetCompany(context.Background(), "kj12sd23-2345-gfds-gv34-gr4egr24ff2d"))
	/*
		fmt.Println(cmpnyService.GetCompany(
			context.Background(), "dafafafafafafa",
		))
	*/
	return nil
}
func main() {
	fmt.Println("Heyyyyyy!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
