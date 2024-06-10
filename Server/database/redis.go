package database

import (
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func GetRedis() *redis.Client {
	return Redis
}
