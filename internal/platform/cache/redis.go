package cache

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func RedisConnection() (*redis.Client, error) {
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))
	options := &redis.Options{
		Addr:     "127.0.0.1:6369",
		Password: "",
		DB:       dbNumber,
	}

	return redis.NewClient(options), nil
}
