package main

import (
	"bee-api-v2/internal/api/router"
	"bee-api-v2/internal/api/server"
	"bee-api-v2/internal/bee"
	"bee-api-v2/internal/config"
	"bee-api-v2/internal/logger"
	"bee-api-v2/internal/requester"
	"flag"
)

func main() {

	var cfgFlag string
	flag.StringVar(&cfgFlag, "c", "config.yaml", "pa")
	// Read config
	cfg := config.ReadConfig(cfgFlag)
	// Created logger
	l := logger.NewLogger()

	// Created requester
	req := requester.NewRequest(cfg)
	// Created core
	app := bee.New(req, cfg, l)
	// Created router
	h := router.NewRouter(app)

	s := server.NewServer(h, cfg)

	s.Start()

}
