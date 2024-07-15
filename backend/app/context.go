package app

import "database/sql"

type Context struct {
	Configuration *Configuration
	Database      *sql.DB
}
