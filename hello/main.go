package main

import f "fmt"

var (
	x, y, z = 16, "Wilson", true
)

func main() {

	v, e := f.Println("Hello World!")
	f.Println(v, e)
	f.Println(x, y, z)
}
