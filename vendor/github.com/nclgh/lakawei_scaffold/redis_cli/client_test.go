package redis_cli

import (
	"testing"
	"fmt"
)

func TestGetRedisClient(t *testing.T) {
	c := GetRedisClient("passport")
	resPing, err := c.Do("ping")
	if err != nil {
		panic(fmt.Sprintf("ping redis server failed. err: %v", err))
	}
	t.Logf("ping redis server success. rsp: %v", resPing.(string))
}
