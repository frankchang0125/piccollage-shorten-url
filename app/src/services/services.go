package services

import (
	"database/sql"

	"github.com/go-redis/redis"
)

var db *sql.DB
var redisClient *redis.Client

func Init(_db *sql.DB, _redisClient *redis.Client) (err error) {
	db = _db
	redisClient = _redisClient
	return InitShortenURLService(db)
}
