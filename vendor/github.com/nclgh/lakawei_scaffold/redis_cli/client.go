package redis_cli

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
)

var RedisAddr = map[string]string{
	"passport": "127.0.0.1:6379",
}

func GetRedisClient(redisServerName string) redis.Conn {
	redisAddr, ok := RedisAddr[redisServerName]
	if !ok {
		panic(fmt.Sprintf("can't find redis server. %v", redisServerName))
	}
	c, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		panic(fmt.Sprintf("connect redis server failed. redis server: %v, err: %v", redisServerName, err))
	}
	resPing, err := c.Do("ping")
	if err != nil {
		panic(fmt.Sprintf("ping redis server failed. redis server: %v, err: %v", redisServerName, err))
	}
	logrus.Infof("ping redis server success. redis server: %v, rsp: %v", redisServerName, resPing.(string))
	return c
}
