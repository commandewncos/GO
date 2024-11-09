package main

import (
	"net/http"
)

type HelloHandler struct{}
type WorldHandler struct{}

func (HelloHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))

}

func (WorldHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("World"))
}

var PORT string = ":5000"

func main() {

	hello := HelloHandler{}
	world := WorldHandler{}
	http.HandleFunc("/hello", hello.ServerHTTP)
	http.HandleFunc("/world", world.ServerHTTP)

	http.ListenAndServe(PORT, nil)

}
