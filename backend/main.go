package main

import (
	"fmt"

	"digitao.io/website/app"
	"digitao.io/website/endpoint"
)

func main() {
	configuration, err := app.ReadConfiguration()
	if err != nil {
		panic(err)
	}

	database, err := app.InitializeDatabase(configuration)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	ctx := app.Context{}
	ctx.Configuration = configuration
	ctx.Database = database

	r := endpoint.SetupRoutes(&ctx)

	r.Run(fmt.Sprintf(":%d", configuration.Port))
}
