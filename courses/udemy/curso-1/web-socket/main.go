package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ls, err := net.Listen("tcp", ":5000")

	if err != nil {
		panic(err)
	}

	defer ls.Close()

	for {
		cnxn, err := ls.Accept()
		if err != nil {
			panic(err)
		}

		go func(cnxn net.Conn) {
			for {
				data, _ := bufio.NewReader(cnxn).ReadString('\n')
				fmt.Println("Conexão iniciada:\t", data)
				if strings.Contains(data, "exit") {
					break
				}
				cnxn.Write([]byte("Sua mensagem foi recebida!"))
			}
			cnxn.Close()
			fmt.Println("Conexão finalizada!")

		}(cnxn)
	}
}
