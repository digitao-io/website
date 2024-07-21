package setup

import (
	"context"
	"fmt"

	"digitao.io/website/app"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDatabase(configuration *app.Configuration) (*mongo.Database, error) {
	opts := options.Client().
		SetHosts(
			[]string{
				fmt.Sprintf(
					"%s:%d",
					configuration.Database.Host,
					configuration.Database.Port,
				),
			},
		).
		SetAuth(
			options.Credential{
				Username: configuration.Database.User,
				Password: configuration.Database.Password,
			},
		).
		SetServerAPIOptions(
			options.ServerAPI(options.ServerAPIVersion1),
		)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	db := client.Database(configuration.Database.Database)

	return db, nil
}
