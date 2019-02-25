package redis

import (
	"fmt"
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/nclgh/lakawei_passport/types"
	"github.com/nclgh/lakawei_scaffold/redis_cli"
)

const SessionExpireTime = 2 * 3600

var passportRedisClient = redis_cli.GetRedisClient("passport")

func GenSessionKey() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

func CreateSession(userId int64) (string, error) {
	sId := GenSessionKey()
	s := &types.Session{
		Key:    sId,
		UserId: userId,
	}
	bV, _ := json.Marshal(s)
	val := string(bV)
	_, err := passportRedisClient.Do("set", sId, val, "ex", SessionExpireTime)
	if err != nil {
		return "", fmt.Errorf("set redis failed. key: %v, err: %v", sId, err)
	}
	return sId, nil
}

func GetSession(sid string) (*types.Session, error) {
	ret, err := passportRedisClient.Do("get", sid)
	if err != nil {
		return nil, err
	}
	s := &types.Session{}
	err = json.Unmarshal(ret.([]byte), s)
	if err != nil {
		return nil, fmt.Errorf("umarshal session err. ret: %v", ret)
	}
	return s, nil
}
