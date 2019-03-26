package main

import (
	"github.com/nclgh/lakawei_passport/handler"
	"github.com/nclgh/lakawei_scaffold/rpc/passport"
)

type ServicePassport struct {
}

func (server *ServicePassport) CreateSession(req passport.CreateSessionRequest, res *passport.CreateSessionResponse) error {
	resp := handler.CreateSession(&req)
	*res = *resp
	return nil
}

func (server *ServicePassport) GetSession(req passport.GetSessionRequest, res *passport.GetSessionResponse) error {
	resp := handler.GetSession(&req)
	*res = *resp
	return nil
}

func (server *ServicePassport) DeleteSession(req passport.DeleteSessionRequest, res *passport.DeleteSessionResponse) error {
	resp := handler.DeleteSession(&req)
	*res = *resp
	return nil
}
