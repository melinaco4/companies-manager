package main

import (
	"fmt"
	"net/http"

	"github.com/melinaco4/companies-manager/internal/company"
	"github.com/melinaco4/companies-manager/internal/database"
	serveHttp "github.com/melinaco4/companies-manager/internal/http"
	log "github.com/sirupsen/logrus"
)

//type App struct{}

func Run() error {
	fmt.Println("Start up app")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		log.Error("failed to setup database")
		return err
	}

	commentService := company.NewService(db)

	handler := serveHttp.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
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
