package main

import (
	"fmt"

	"github.com/melinaco4/companies-manager/internal/company"
	"github.com/melinaco4/companies-manager/internal/database"
	serveHttp "github.com/melinaco4/companies-manager/internal/http"
)

//type App struct{}

func Run() error {
	fmt.Println("starting up our application")

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmpnService := company.NewService(db)
	/*
		fmt.Println(cmpnService.PostCompany(context.Background(),
			company.Company{
				ID:                "c2c49920-682b-11ed-8c0b-1564f1b782d6",
				Name:              "Twitter",
				Description:       "blabla",
				AmountofEmployees: 13,
				Registered:        true,
				Type:              "test",
			}))
	*/

	httpHandler := serveHttp.NewHandler(cmpnService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}
func main() {
	fmt.Println("Heyyyyyy!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
