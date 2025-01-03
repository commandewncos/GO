package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Running...")
	log.Fatal(http.ListenAndServe(":5500", nil))
}
