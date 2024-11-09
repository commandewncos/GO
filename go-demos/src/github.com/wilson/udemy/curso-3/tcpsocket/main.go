package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":5500")

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		fmt.Println("Conexão estabelecida.")
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			for {
				data, _ := bufio.NewReader(conn).ReadString('\n')
				fmt.Println("Dado recebido:", data)
				if data == "exit" {
					break
				}
				conn.Write([]byte("Sua mensagem foi recebida com sucesso!\n"))

			}

			conn.Close()
			fmt.Println("Conexão encerrada.")
		}(conn)
	}

}
