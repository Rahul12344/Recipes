package redis

import (
	"github.com/go-redis/redis"
)

// Connect connects to redis
func Connect(host, port, password string, db int) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	return redisClient
}
