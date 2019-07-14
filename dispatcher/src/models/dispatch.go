package models

import (
	"database/sql"
	"fmt"
)

type DispatchDao interface {
	GetCounter() (counter uint64, err error)
	UpdateCounter(counter uint64) (err error)
}

type DispatchSQLDao struct {
	db *sql.DB
}

func NewDispatchSQLDao(db *sql.DB) *DispatchSQLDao {
	return &DispatchSQLDao{
		db: db,
	}
}

func (c *DispatchSQLDao) GetCounter() (counter uint64, err error) {
	query := fmt.Sprintf(`
		SELECT counter
		FROM %s
		WHERE id = 1`, dispatchTable)
	err = c.db.QueryRow(query).Scan(&counter)
	return
}

func (c *DispatchSQLDao) UpdateCounter(counter uint64) (err error) {
	query := fmt.Sprintf(`
		UPDATE %s
		SET counter = ?
		WHERE id = 1`, dispatchTable)
	_, err = c.db.Exec(query, counter)
	return
}
