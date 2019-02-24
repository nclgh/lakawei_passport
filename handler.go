package main

import (
	"github.com/koding/kite"
	"lakawei/lakawei_scaffold/kite/passport"
	"lakawei/lakawei_scaffold/kite/kite_common"
	"github.com/sirupsen/logrus"
)

func CreateSession(r *kite.Request) (interface{}, error) {
	str, err := r.Args.MarshalJSON()
	if err != nil {
		logrus.Errorf("err: %v", err)
		return passport.GetCreateSessionResponse(kite_common.CodeReqErr), nil
	}
	str2 := string(str)
	logrus.Infof("str: %v", str2)
	req, err := passport.UnmarshalCreateSessionRequest(str2)
	if err != nil {
		logrus.Infof("err: %v", err)
		return passport.GetCreateSessionResponse(kite_common.CodeReqErr), nil
	} else {
		logrus.Infof("req: %v", *req)
	}
	return passport.GetCreateSessionResponse(kite_common.CodeSuccess), nil // send back the result
}
