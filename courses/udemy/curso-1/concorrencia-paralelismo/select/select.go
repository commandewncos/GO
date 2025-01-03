package main

import (
	"fmt"
	"time"

	h "github.com/Wilson-cmd/html"
)

func ohMaisRapido(url1, url2, url3 string) string {
	c1 := h.Titulo(url1)
	c2 := h.Titulo(url2)
	c3 := h.Titulo(url3)

	select {
	case t1 := <-c1:
		return t1

	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3

	case <-time.After(time.Millisecond * 1000):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta"
	}

}

func main() {
	campeao := ohMaisRapido(
		"https://alura.com.br",
		"https://google.com.br",
		"https://youtube.com.br",
	)

	fmt.Println(campeao)

}
