package main

import (
	"fmt"
	"runtime"
)

func main() {
	re := runtime.NumCPU()

	fmt.Println(re)
}
