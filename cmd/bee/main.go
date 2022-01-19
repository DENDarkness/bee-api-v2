package main

import (
	"bee-api-v2/internal/config"
	"bee-api-v2/internal/launcher"
	"bee-api-v2/internal/logger"
	"flag"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {

	var cfgFlag string
	flag.StringVar(&cfgFlag, "c", "config.yaml", "pa")

	// Read config
	cfg := config.ReadConfig(cfgFlag)

	// Created logger
	l := logger.NewLogger()

	// Create memory cache
	c := cache.New(5*time.Second, 10*time.Second)

	app := launcher.NewApp(cfg, l, c)

	app.Launch()

}
