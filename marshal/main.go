package main

import (
	"encoding/json"
	"log"
)

type Pessoa struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	p := []Pessoa{
		{Name: "Wilson-1", Age: 35},
		{Name: "Wilson-2", Age: 36},
		{Name: "Wilson-2", Age: 37},
		{Name: "Wilson-3", Age: 38},
		{Name: "Wilson-4", Age: 39},
		{Name: "Wilson-5", Age: 31},
		{Name: "Wilson-6", Age: 33},
		{Name: "Wilson-7", Age: 32},
	}

	s := []byte(`[{"name": "wilson", "age": 31}]`)

	if e := json.Unmarshal(s, &p); e != nil {
		log.Println(e)
	}

}
