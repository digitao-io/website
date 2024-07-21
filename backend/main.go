package main

import (
	"context"
	"fmt"

	"digitao.io/website/app"
	"digitao.io/website/dataendpoint"
	"digitao.io/website/setup"
	"digitao.io/website/siteendpoint"
	"github.com/gin-gonic/gin"
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
	defer (func() {
		err := database.Client().Disconnect(context.Background())
		if err != nil {
			panic(err)
		}
	})()

	objstorage, err := setup.SetupObjstorage(configuration)
	if err != nil {
		panic(err)
	}

	ctx := app.Context{}
	ctx.Configuration = configuration
	ctx.Database = database
	ctx.Objstorage = objstorage

	r := gin.Default()

	siteendpoint.SetupRoutes(&ctx, r)
	dataendpoint.SetupRoutes(&ctx, r)

	r.Run(fmt.Sprintf(":%d", configuration.Port))
}
