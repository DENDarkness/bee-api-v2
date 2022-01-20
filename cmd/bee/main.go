package main

import (
	"bee-api-v2/internal/config"
	"bee-api-v2/internal/launcher"
	"flag"
)

func main() {

	var cfgFlag string
	flag.StringVar(&cfgFlag, "c", "config.yaml", "path to config file")

	// Read config
	cfg := config.ReadConfig(cfgFlag)

	app := launcher.NewApp(cfg)

	app.Launch()

}
