package redis_cli

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisAddr = map[string]string{
	"passport": "127.0.0.1:6379",
}

func GetRedisClient(redisServerName string) *redis.Client {
	redisAddr, ok := RedisAddr[redisServerName]
	if !ok {
		panic(fmt.Sprintf("can't find redis server. %v", redisServerName))
	}
	c := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	err := c.Ping().Err()
	if err != nil {
		panic(fmt.Sprintf("ping redis server failed. redis server: %v, err: %v", redisServerName, err))
	}
	return c
}
