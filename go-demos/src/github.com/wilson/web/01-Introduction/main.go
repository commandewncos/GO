package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloFunction(w http.ResponseWriter, r *http.Request) {
	// Using constants at header
	w.WriteHeader(http.StatusOK) // -> status: 200
	w.Write([]byte("Hello world!"))
	fmt.Println("Server running...")
}

func homeFunction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome!"))

}

type API struct{}

func (a API) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello again..."))

}

func main() {

	PORT := ":3000"

	svr := API{}

	http.HandleFunc("/", homeFunction)
	http.HandleFunc("/hello", helloFunction)
	http.HandleFunc("/api", svr.ServerHTTP)
	// // When have error
	// err := http.ListenAndServe(":5000", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Other way to trate error with log.Fatal
	log.Fatal(http.ListenAndServe(PORT, nil))

}
