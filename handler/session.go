package handler

import (
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_passport/redis"
	"github.com/nclgh/lakawei_scaffold/utils"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
	"github.com/nclgh/lakawei_scaffold/rpc/passport"
)

func CreateSession(req *passport.CreateSessionRequest) (rsp *passport.CreateSessionResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("CreateSession panic: %v, stack: %v", err, stacks)
		rsp = getCreateSessionResponse(common.CodeFailed, "panic")
	})
	sId, err := redis.CreateSession(req.UserId)
	if err != nil {
		logrus.Errorf("create session err: %v", err)
		return getCreateSessionResponse(common.CodeFailed, "")
	}
	rsp = getCreateSessionResponse(common.CodeSuccess, "")
	rsp.SessionId = sId
	return rsp
}

func getCreateSessionResponse(code common.RspCode, msg string) *passport.CreateSessionResponse {
	rsp := &passport.CreateSessionResponse{
		SessionId: "",
		Code:      code,
		Msg:       msg,
	}
	return rsp
}

func GetSession(req *passport.GetSessionRequest) (rsp *passport.GetSessionResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("GetSession panic: %v, stack: %v", err, stacks)
		rsp = getGetSessionResponse(common.CodeFailed, "panic")
	})
	s, err := redis.GetSession(req.SessionId)
	if err != nil {
		logrus.Errorf("get session err: %v", err)
		return getGetSessionResponse(common.CodeFailed, "")
	}
	rsp = getGetSessionResponse(common.CodeSuccess, "")
	if s != nil {
		rsp.UserId = s.UserId
	} else {
		rsp.UserId = 0
	}
	return rsp
}

func getGetSessionResponse(code common.RspCode, msg string) *passport.GetSessionResponse {
	rsp := &passport.GetSessionResponse{
		UserId: 0,
		Code:   code,
		Msg:    msg,
	}
	return rsp
}

func DeleteSession(req *passport.DeleteSessionRequest) (rsp *passport.DeleteSessionResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("DeleteSession panic: %v, stack: %v", err, stacks)
		rsp = getDeleteSessionResponse(common.CodeFailed, "panic")
	})
	err := redis.DeleteSession(req.UserId)
	if err != nil {
		logrus.Errorf("delete session err: %v", err)
		return getDeleteSessionResponse(common.CodeFailed, "")
	}
	rsp = getDeleteSessionResponse(common.CodeSuccess, "")
	rsp.IsSuccess = true
	return rsp
}

func getDeleteSessionResponse(code common.RspCode, msg string) *passport.DeleteSessionResponse {
	rsp := &passport.DeleteSessionResponse{
		IsSuccess: false,
		Code:      code,
		Msg:       msg,
	}
	return rsp
}
