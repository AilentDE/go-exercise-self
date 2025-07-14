package cache

import "github.com/redis/go-redis/v9"

var (
	RedisClient *redis.Client
)

func InitRedisClient(addr string) *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	return RedisClient
}
