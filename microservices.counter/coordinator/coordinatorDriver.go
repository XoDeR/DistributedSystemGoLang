package main

import (
	"microservices.counter/common"

	"github.com/NYTimes/gizmo/config"
	"github.com/NYTimes/gizmo/server"
)

func main() {
	var cfg *common.Config
	config.LoadJSONFile("./../config.json", &cfg)
	server.SetConfigOverrides(cfg.Server)
	server.Init("Microservices.Counter", cfg.Server)

	err := server.Register(common.NewRPCService(cfg))
	if err != nil {
		server.Log.Fatal("unable to register service: ", err)
	}

	err1 := server.Run()
	if err1 != nil {
		server.Log.Fatal("server encountered a fatal error: ", err1)
	}
}
