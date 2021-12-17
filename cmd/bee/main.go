package main

import (
	"bee-api-v2/internal/api/router"
	"bee-api-v2/internal/api/server"
	"bee-api-v2/internal/bee"
	"bee-api-v2/internal/config"
	"bee-api-v2/internal/requester"
	"flag"
)

func main() {

	var cfgFlag string
	flag.StringVar(&cfgFlag, "c", "config.yaml", "path configuration file")

	cfg := config.ReadConfig(cfgFlag)

	r := requester.NewRequest()
	app := bee.New(r)
	h := router.NewRouter(*app)

	s := server.NewServer(h, cfg)

	s.Start()

}
