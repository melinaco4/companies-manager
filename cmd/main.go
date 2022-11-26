package main

import (
	"fmt"

	"github.com/melinaco4/companies-manager/internal/run"
)

func main() {
	fmt.Println("Hello!")
	app := run.App{}
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
