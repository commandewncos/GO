package main

import (
	"fmt"
	"runtime"
	"sync"
)

const totalGoRoutines int = 10

var counter int = 0

var wg sync.WaitGroup

var mu sync.Mutex

func main() {

	wg.Add(totalGoRoutines)

	for i := 0; i <= totalGoRoutines; i++ {
		go func() {
			v := counter
			mu.Lock()
			// yield
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()

		}()
	}

	wg.Wait()
	fmt.Println(counter)

}
