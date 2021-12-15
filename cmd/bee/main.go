package main

import (
	"bee-api-v2/internal/api/router"
	"bee-api-v2/internal/api/server"
	"bee-api-v2/internal/bee"
	"bee-api-v2/internal/requester"
)

func main() {



	r := requester.NewRequest()
	app := bee.New(r)
	h := router.NewRouter(*app)

	s := server.NewServer(h)

	s.Start()

}
