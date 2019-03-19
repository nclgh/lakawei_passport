package main

import "github.com/nclgh/lakawei_rpc/server"

func main() {
	server.Init()
	server.Run(new(ServicePassport))
}
