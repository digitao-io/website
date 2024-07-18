package main

import (
	"fmt"

	"digitao.io/website/app"
	"digitao.io/website/setup"
)

func main() {
	configuration, err := setup.ReadConfiguration()
	if err != nil {
		panic(err)
	}

	database, err := setup.SetupDatabase(configuration)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	sqlBuilder := setup.SetupSqlBuilder()

	objstorage, err := setup.SetupObjstorage(configuration)
	if err != nil {
		panic(err)
	}

	ctx := app.Context{}
	ctx.Configuration = configuration
	ctx.Database = database
	ctx.SqlBuilder = sqlBuilder
	ctx.Objstorage = objstorage

	r := setup.SetupRoutes(&ctx)

	r.Run(fmt.Sprintf(":%d", configuration.Port))
}
