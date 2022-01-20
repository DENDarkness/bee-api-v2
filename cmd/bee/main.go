package main

import (
	"bee-api-v2/internal/config"
	"bee-api-v2/internal/launcher"
	"flag"
)

func main() {

	var cfgFlag string
	flag.StringVar(&cfgFlag, "c", "config.yaml", "pa")

	// Read config
	cfg := config.ReadConfig(cfgFlag)

	app := launcher.NewApp(cfg)

	app.Launch()

}
