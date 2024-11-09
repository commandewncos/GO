package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const n = 5
const logFileName = "log.txt"
const fileName = "source.txt"

func main() {

	urls := readFile(fileName)
	for {
		welcomeProgram()
		switch selectChoice() {
		case 0:
			fmt.Println("")
			os.Exit(0)
		case 1:
			watching(urls)
		case 2:
			returnLog()
		default:
			fmt.Println("Param don't exist!")
			os.Exit(-1)
		}
	}

}

// FUNCTIONS

func welcomeProgram() {
	fmt.Println("Option 0: exit")
	fmt.Println("Option 1: Start the program")
	fmt.Println("Option 2: Show logs")
}

func selectChoice() int {
	var choice int
	fmt.Scan(&choice)
	return choice

}

func watching(param []string) {
	for h := 0; h <= n; h++ {

		for _, i := range param {
			verify(i)
		}

		time.Sleep(n * time.Second)

	}

}

// Verify status code of website
func verify(link string) {
	response, _ := http.Get(link)
	if response.StatusCode == 200 {
		fmt.Println(link, "-> Status code: 200")
		logMonitor(link, true)
	} else {
		fmt.Println(response.StatusCode)
		logMonitor(link, false)
	}
}

func readFile(param string) []string {
	var list []string
	file, err := os.Open(param)
	if err != nil {
		fmt.Println("Deu erro meu mano!", err)

	} else {
		read := bufio.NewReader(file)
		for {
			r, err := read.ReadString('\n')
			r = strings.TrimSpace(r)
			list = append(list, r)
			if err == io.EOF {
				break
			}
		}

	}

	file.Close()
	return list

}

func logMonitor(site string, status bool) {
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(time.Now().Format("02/01/2006 15:04:05") + "\t" + site + "\tonline:\t" + strconv.FormatBool(status) + "\n")
	}

	file.Close()

}

func returnLog() {
	file, err := os.ReadFile(logFileName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))

}
