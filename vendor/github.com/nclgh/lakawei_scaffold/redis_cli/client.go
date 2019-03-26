package redis_cli

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/nclgh/lakawei_scaffold/conf"
)

var (
	clients = make(map[string]*redis.Client)
)

func GetRedisClient(redisServerName string) *redis.Client {
	c := conf.GetConfig()
	redisAddr := c.DefaultString(fmt.Sprintf("%sRedisAddr", redisServerName), "")
	if redisAddr == "" {
		panic(fmt.Sprintf("can't find redis server. %v", redisServerName))
	}
	if cli, ok := clients[redisServerName]; ok {
		return cli
	}
	cli := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	err := cli.Ping().Err()
	if err != nil {
		panic(fmt.Sprintf("ping redis server failed. redis server: %v, err: %v", redisServerName, err))
	}
	clients[redisServerName] = cli
	return cli
}
