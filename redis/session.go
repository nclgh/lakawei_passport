package redis

import (
	"fmt"
	"time"
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/go-redis/redis"
	"github.com/nclgh/lakawei_passport/types"
	"github.com/nclgh/lakawei_scaffold/redis_cli"
	"github.com/sirupsen/logrus"
)

const SessionExpireTime = 2 * time.Hour

var passportRedisClient = redis_cli.GetRedisClient("passport")

func GenSessionKey() string {
	uid := uuid.NewV4()
	return uid.String()
}

func getSessionRedisKey(sid string) string {
	return fmt.Sprintf("passport_sid_%s", sid)
}

func getUserIdRedisKey(userId int64) string {
	return fmt.Sprintf("passport_userid_%v", userId)
}

func CreateSession(userId int64) (string, error) {
	// 先登出 把原来的session干掉
	ret, err := passportRedisClient.Get(getUserIdRedisKey(userId)).Result()
	if err != nil {
		if err != redis.Nil {
			return "", fmt.Errorf("get old session from redis err: %v", err)
		}
	} else {
		err := passportRedisClient.Del(getSessionRedisKey(ret)).Err()
		if err != nil {
			return "", fmt.Errorf("del old session from redis err: %v", err)
		}
	}

	// 创建新的session
	sId := GenSessionKey()

	err = passportRedisClient.Set(getUserIdRedisKey(userId), sId, SessionExpireTime).Err()
	if err != nil {
		return "", fmt.Errorf("set redis userId to session err: %v", err)
	}

	s := &types.Session{
		Key:    sId,
		UserId: userId,
	}
	bV, _ := json.Marshal(s)
	err = passportRedisClient.Set(getSessionRedisKey(sId), string(bV), SessionExpireTime).Err()
	if err != nil {
		return "", fmt.Errorf("set redis failed. key: %v, err: %v", sId, err)
	}
	return sId, nil
}

func GetSession(sId string) (*types.Session, error) {
	ret, err := passportRedisClient.Get(getSessionRedisKey(sId)).Result()
	if err != nil {
		return nil, err
	}
	s := &types.Session{}
	err = json.Unmarshal([]byte(ret), s)
	if err != nil {
		return nil, fmt.Errorf("umarshal session err. ret: %v", ret)
	}
	return s, nil
}

func DeleteSession(userId int64) error {
	sid, err := passportRedisClient.Get(getUserIdRedisKey(userId)).Result()
	if err != nil {
		return fmt.Errorf("delete session get sid err: %v", err)
	}
	err = passportRedisClient.Del(getSessionRedisKey(sid)).Err()
	if err != nil {
		return fmt.Errorf("delete session del session err: %v", err)
	}
	// 删除userId-sid的key 失败不用返回错误
	err = passportRedisClient.Del(getUserIdRedisKey(userId)).Err()
	if err != nil {
		logrus.Warnf("delete session del userId-session err: %v", err)
	}
	return nil
}
