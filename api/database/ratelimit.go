package database

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

// Checks the rate limit of the user based on their ip
func CheckRateLimit(ip string) gin.H {
	_, err := RateClient.Get(Ctx, ip).Result()

	// Checking if user is registering for first time
	if err == redis.Nil {
		RateClient.Set(Ctx, ip, os.Getenv("API_QUOTA"), 30*time.Minute)
	}

	// Getting the requests after initiating
	val, _ := RateClient.Get(Ctx, ip).Result()
	valInt, _ := strconv.Atoi(val)

	// Checking for remaining requests
	if valInt <= 0 {
		limit, _ := RateClient.TTL(Ctx, ip).Result()
		return gin.H{
			"error":           "rate limit exceeded",
			"rate_limit_rest": limit / time.Nanosecond / time.Minute,
		}
	}

	return nil
}

// Decreasing remaining requests
func DecrRateLimit(ip string) int {
	rl, _ := RateClient.Decr(Ctx, ip).Result()
	return int(rl)
}
