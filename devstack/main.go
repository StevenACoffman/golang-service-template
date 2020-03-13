package main

import (
	"github.com/pseudo-su/golang-service-template/internal"
	"github.com/pseudo-su/golang-service-template/internal/config"
)

func main() {
	cfg := config.NewApplicationConfig()
	server := internal.Bootstrap(cfg)

	server.Start(cfg.ServerPort())
}
