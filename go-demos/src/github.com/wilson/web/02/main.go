package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
Server static content with each function

func contentHome(w http.ResponseWriter, r *http.Request) {

	file, err := os.ReadFile("static/home.html")

	if err != nil {
		log.Fatal(err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(file)

	}

}

func contentAbout(w http.ResponseWriter, r *http.Request) {
	f, e := os.ReadFile("static/about.html")

	if e != nil {
		log.Fatal(e)
	} else {
		w.Write(f)
	}
}

*/

func StaticHandler(w http.ResponseWriter, r *http.Request) {

	request := r.URL.Path

	f, err := os.Open("static" + request)

	if err != nil {
		log.Println(err)
		return
	}

	if strings.HasSuffix(r.URL.Path, ".css") {
		w.Header().Add("Content-Type", "text/css")
	}

	io.Copy(w, f)
}
func main() {
	PORT := ":5000"
	// http.HandleFunc("/", contentHome)
	// http.HandleFunc("/about", contentAbout)

	http.HandleFunc("/", StaticHandler)

	http.ListenAndServe(PORT, nil)
}
