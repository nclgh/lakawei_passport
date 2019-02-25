package main

import (
	"github.com/koding/kite"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_passport/helper"
	"github.com/nclgh/lakawei_passport/handler"
	"github.com/nclgh/lakawei_scaffold/kite/passport"
	"github.com/nclgh/lakawei_scaffold/kite/kite_common"
)

func CreateSession(r *kite.Request) (rsp *passport.CreateSessionResponse, rErr error) {
	defer helper.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("CreateSession panic: %v, stack: %v", err, stacks)
		rsp = passport.GetCreateSessionResponse(kite_common.CodeFailed, "panic")
	})
	str, err := r.Args.MarshalJSON()
	if err != nil {
		logrus.Errorf("err: %v", err)
		return passport.GetCreateSessionResponse(kite_common.CodeReqErr, ""), nil
	}
	req, err := passport.UnmarshalCreateSessionRequest(string(str))
	if err != nil {
		logrus.Infof("err: %v", err)
		return passport.GetCreateSessionResponse(kite_common.CodeReqErr, ""), nil
	}
	//logrus.Infof("req: %v", *req)
	rsp, err = handler.CreateSession(req)
	if err != nil {
		return passport.GetCreateSessionResponse(kite_common.CodeFailed, ""), nil
	}
	return rsp, nil
}

func GetSession(r *kite.Request) (rsp *passport.GetSessionResponse, rErr error) {
	defer helper.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("GetSession panic: %v, stack: %v", err, stacks)
		rsp = passport.GetGetSessionResponse(kite_common.CodeFailed, "panic")
	})
	str, err := r.Args.MarshalJSON()
	if err != nil {
		logrus.Errorf("err: %v", err)
		return passport.GetGetSessionResponse(kite_common.CodeReqErr, ""), nil
	}
	req, err := passport.UnmarshalGetSessionRequest(string(str))
	if err != nil {
		logrus.Infof("err: %v", err)
		return passport.GetGetSessionResponse(kite_common.CodeReqErr, ""), nil
	}
	//logrus.Infof("req: %v", *req)
	rsp, err = handler.GetSession(req)
	if err != nil {
		return passport.GetGetSessionResponse(kite_common.CodeFailed, ""), nil
	}
	return rsp, nil
}
