package run

import (
	"fmt"

	"github.com/melinaco4/companies-manager/internal/company"
	"github.com/melinaco4/companies-manager/internal/database"
	serveHttp "github.com/melinaco4/companies-manager/internal/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Application started and ready to go!")

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmpnService := company.NewService(db)

	httpHandler := serveHttp.NewHandler(cmpnService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}
