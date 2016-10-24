package main

import (
	"microservices.counter/service"

	"github.com/NYTimes/gizmo/config"
	"github.com/NYTimes/gizmo/server"
)

func main() {
	// showing 1 way of managing gizmo/config: importing from a local file
	var cfg *service.Config
	config.LoadJSONFile("./config.json", &cfg)
	server.SetConfigOverrides(cfg.Server)

	server.Init("nyt-rpc-proxy", cfg.Server)

	err := server.Register(service.NewRPCService(cfg))
	if err != nil {
		server.Log.Fatal("unable to register service: ", err)
	}

	err = server.Run()
	if err != nil {
		server.Log.Fatal("server encountered a fatal error: ", err)
	}
}
