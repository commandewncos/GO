package main

import (
	"go/api/database"
	"go/api/routes"
)

func main() {

	database.ConnectDB()
	routes.HandleAllRequests()

}
