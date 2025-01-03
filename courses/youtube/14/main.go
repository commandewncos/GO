package main

import (
	f "fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	f.Fprintf(w, "<h1>Hello World!</h1>")

}

func about(w http.ResponseWriter, r *http.Request) {

	f.Fprintf(w, "<h1>About something...</h1>")

}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	f.Println("Server Starting...")

	http.ListenAndServe(":3000", nil)

}
