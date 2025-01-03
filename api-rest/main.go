package main

import (
	"rest-api/connection"
	"rest-api/routes"
)

func main() {
	connection.ConnectionWithDatabase()

	routes.HandleRequests()

}
