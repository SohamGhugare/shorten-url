package database

import (
	"context"
	"errors"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Client *redis.Client     // actual database client with all the urls
var RateClient *redis.Client // another client for storing rate limits

// Creates a new redis client based on your environment variables
func CreateUrlClient(dbNo int) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	Client = rdb
}

// Creating a Rate limit client
func CreateRateClient(dbNo int) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	RateClient = rdb
}

// Create url
func CreateURL(url string, short string) error {
	err := Client.Set(Ctx, short, url, 0).Err()
	return err
}

// Finding an url
func GetURL(short string) (string, error) {
	val, err := Client.Get(Ctx, short).Result()
	if err == redis.Nil {
		return "", errors.New("url not found")
	}
	return val, err
}
