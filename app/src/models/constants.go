package models

import (
	"errors"
)

// Database
const redirectsTable = "redirects"

var (
	ErrShortenURLNotExists = errors.New("Shorten URL does not exists")
	ErrURLNotExists        = errors.New("URL does not exists")
	ErrDuplicateURL        = errors.New("Duplicate URL")
)
