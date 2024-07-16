package setup

import (
	"database/sql"
	"fmt"

	"digitao.io/website/app"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/go-sql-driver/mysql"
)

func SetupDatabase(configuration *app.Configuration) (*sql.DB, error) {
	dbConfig := mysql.NewConfig()

	dbConfig.Addr = fmt.Sprintf("%s:%d", configuration.Database.Host, configuration.Database.Port)
	dbConfig.User = configuration.Database.User
	dbConfig.Passwd = configuration.Database.Password
	dbConfig.DBName = configuration.Database.Database

	database, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}

	database.SetConnMaxIdleTime(0)
	database.SetConnMaxLifetime(0)
	database.SetMaxIdleConns(configuration.Database.MaxIdleConnections)
	database.SetMaxOpenConns(configuration.Database.MaxOpenConnections)

	return database, nil
}

func SetupSqlBuilder() *goqu.DialectWrapper {
	sqlBuilder := goqu.Dialect("mysql")
	return &sqlBuilder
}
