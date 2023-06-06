package main

import (
	"github.com/dadadam/go-http-example/api/server"
	"github.com/dadadam/go-http-example/configs"
)

func main() {

	configs.Init()
	cfg := configs.GetConfig()

	server := server.NewServer(cfg)
	server.Start()
}
