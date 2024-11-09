package main

import (
	"net/http"
)

func main() {

	PORT := ":5500"

	// static := http.Dir("static")
	// staticHandler := http.FileServer(static)

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(PORT, nil)
}
