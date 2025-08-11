package main

import (
	"log"

	"github.com/lakshsetia/jwt-authentication/internal/config"
	"github.com/lakshsetia/jwt-authentication/internal/server"
)

func main() {
	// load config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// create new app
	app, err := server.NewApp(config)
	if err != nil {
		log.Fatal(err)
	}
	// start app
	app.Run()
}