package services

import (
	"database/sql"
)

var db *sql.DB

func Init(_db *sql.DB) (err error) {
	db = _db
	return InitDispatchService(db)
}
