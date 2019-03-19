package main

import (
	"github.com/nclgh/lakawei_passport/handler"
	"github.com/nclgh/lakawei_scaffold/rpc/passport"
)

type ServicePassport struct {
}

func (server *ServicePassport) CreateSession(req passport.CreateSessionRequest, res *passport.CreateSessionResponse) error {
	resp, _ := handler.CreateSession(&req)
	*res = *resp
	return nil
}

func (server *ServicePassport) GetSession(req passport.GetSessionRequest, res *passport.GetSessionResponse) error {
	resp, _ := handler.GetSession(&req)
	*res = *resp
	return nil
}
