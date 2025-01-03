package funcs

import (
	"fmt"
	"time"
)

type functions struct {
}

var Functions functions

func (f *functions) GetYear() int {
	return time.Now().Year()
}

func (f *functions) GetDate() string {
	currentTime := time.Now()
	return fmt.Sprintf("%d-%d-%d", currentTime.Day(), currentTime.Month(), currentTime.Year())

}
