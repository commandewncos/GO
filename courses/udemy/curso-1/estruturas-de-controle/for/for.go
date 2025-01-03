package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0

	for i <= 100 {
		if i%2 == 0 {
			fmt.Println(i)
			time.Sleep(time.Second)

		}
		i++

	}

}
