package main

import (
	"github.com/auth/command/database/mysql/connection"
	"github.com/auth/command/router"
)

func main() {

	connection.ConnectionWithDatabase()

	router.HandleRequest()

}
