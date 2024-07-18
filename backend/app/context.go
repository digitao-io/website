package app

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/minio/minio-go/v7"
)

type Context struct {
	Configuration *Configuration
	Database      *sql.DB
	SqlBuilder    *goqu.DialectWrapper
	Objstorage    *minio.Client
}
