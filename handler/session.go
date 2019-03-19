package handler

import (
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_passport/redis"
	"github.com/nclgh/lakawei_scaffold/rpc/passport"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

func CreateSession(req *passport.CreateSessionRequest) (*passport.CreateSessionResponse, error) {
	sId, err := redis.CreateSession(req.UserId)
	if err != nil {
		logrus.Errorf("create session err: %v", err)
		return passport.GetCreateSessionResponse(common.CodeFailed, ""), nil
	}
	rsp := passport.GetCreateSessionResponse(common.CodeSuccess, "")
	rsp.SessionId = sId
	return rsp, nil
}

func GetSession(req *passport.GetSessionRequest) (*passport.GetSessionResponse, error) {
	s, err := redis.GetSession(req.SessionId)
	if err != nil {
		logrus.Errorf("get session err: %v", err)
		return passport.GetGetSessionResponse(common.CodeFailed, ""), nil
	}
	rsp := passport.GetGetSessionResponse(common.CodeSuccess, "")
	if s != nil {
		rsp.UserId = s.UserId
	} else {
		rsp.UserId = 0
	}
	return rsp, nil
}
