package postgres

import "database/sql"

type Storage struct {
	db *sql.DB
}
