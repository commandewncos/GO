package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.SetOutput(file)
	}

	var digit int
	//Matches a non-digit character.
	cpf := regexp.MustCompile(`\D`).ReplaceAllString("410.060.868-00", "")
	//Length's cpf
	length := len(cpf)
	//Split first (9) digits
	re := strings.Split(cpf[:(length-2)], "")

	for i := (length - 1); i <= length; i++ {
		for index, value := range re {
			converted, err := strconv.Atoi(value)
			if err != nil {
				log.Println(err.Error())
				log.SetOutput(file)
			}
			digit += converted * (i - index)
		}
		recovered := strconv.Itoa(calc(digit, length))
		digit = 0
		re = append(re, recovered)

		fmt.Println(re)
	}

}

// Calculate the digit
func calc(x, y int) int {

	return y - (x % y)
}
