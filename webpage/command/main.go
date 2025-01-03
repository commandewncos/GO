package main

import (
	"log"

	"github.com/Wison-cmd/GO/command/config"
)

func main() {
	app := config.Application{}
	app.Config.CommandsFlag()

	//Start server
	log.Printf("Version: %s | Server %s listenning on port %s", app.Config.Version, app.Config.Env, app.Config.Port)
	log.Fatal(app.Start())

}
