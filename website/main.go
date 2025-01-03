package main

import (
	"net/http"

	"github.com/Wilson-cmd/GO/connection"
	"github.com/Wilson-cmd/GO/router"
)

func main() {
	connection.ConnectionWithDatabase()
	router.GetRoutes()
	http.ListenAndServe(":8081", nil)
}
