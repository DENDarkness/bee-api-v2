package main

import (
	"bee-api-v2/internal/bee"
	"bee-api-v2/internal/requester"
)

func main() {

	r := requester.NewRequest()
	app := bee.New(r)

	app.ModemReboot()

}
