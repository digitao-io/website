package app

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type Context struct {
	Configuration *Configuration
	Database      *sql.DB
	SqlBuilder    *goqu.DialectWrapper
}
