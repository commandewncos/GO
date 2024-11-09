package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func serverInfo(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"int":    1,
		"string": "Wilson",
		"Male":   true,
		"float":  1.2,
		"time":   time.Now(),
		"complex": map[string]interface{}{
			"array": []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},

		"nil": nil,
	}

	response, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

// Working with JSON
func main() {

	http.HandleFunc("/", serverInfo)
	http.ListenAndServe(":3000", nil)
}
