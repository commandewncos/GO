package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, `<body style="color: white; background-color: black;">
					<h1>Hora certa: %s</h1>
					</body>`, s)
}

func main() {

	http.HandleFunc("/horaCerta", horaCerta)

	log.Println("Running...")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
