package app

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func InitializeDatabase(configuration *Configuration) (*sql.DB, error) {
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
