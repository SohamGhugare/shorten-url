package database

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var RateClient *redis.Client // Client for storing rate limits

// Creating a Rate limit client
func CreateRateClient(dbNo int) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	RateClient = rdb
}
