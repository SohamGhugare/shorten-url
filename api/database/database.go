package database

import (
	"os"

	"github.com/redis/go-redis/v9"
)

// Creates a new redis client based on your environment variables
func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return rdb
}
