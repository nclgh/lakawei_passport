package main

import (
	"github.com/koding/kite"
	"lakawei/lakawei_scaffold/kite/kite_common"
)

func main() {
	serverName := "passport"
	k := kite.New(serverName, "1.0.0")

	SetFuncRouter(k)

	k.Config.Port = kite_common.GetServerPort(serverName)
	k.Run()
}

func SetFuncRouter(k *kite.Kite) {
	k.HandleFunc("CreateSession", func(r *kite.Request) (interface{}, error) {
		return CreateSession(r)
	}).DisableAuthentication()
}
