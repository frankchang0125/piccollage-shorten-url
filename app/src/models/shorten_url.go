package models

import (
	"database/sql"
	"fmt"
	"strings"
)

type ShortenURLDao interface {
	CreateShortenURL(url string, shorten string) (err error)
	GetURL(shorten string) (url string, err error)
	GetShortenURL(url string) (shorten string, err error)
}

type ShortenURLSQLDao struct {
	db *sql.DB
}

func NewShortenURLSQLDao(db *sql.DB) *ShortenURLSQLDao {
	return &ShortenURLSQLDao{
		db: db,
	}
}

func (c *ShortenURLSQLDao) CreateShortenURL(url string, shorten string) (err error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (url, shorten)
		VALUES (?, ?)`, redirectsTable)
	_, err = c.db.Exec(query, url, shorten)
	if err != nil {
		if strings.Contains(err.Error(), "") {
			return ErrDuplicateURL
		}
		return err
	}
	return nil
}

func (c *ShortenURLSQLDao) GetURL(shorten string) (url string, err error) {
	query := fmt.Sprintf(`
		SELECT url
		FROM %s
		WHERE shorten = ?`, redirectsTable)
	err = c.db.QueryRow(query, shorten).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrShortenURLNotExists
			return
		}
		return
	}
	return
}

func (c *ShortenURLSQLDao) GetShortenURL(url string) (shorten string, err error) {
	query := fmt.Sprintf(`
		SELECT shorten
		FROM %s
		WHERE url = ?`, redirectsTable)
	err = c.db.QueryRow(query, url).Scan(&shorten)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrURLNotExists
			return
		}
		return
	}
	return
}
