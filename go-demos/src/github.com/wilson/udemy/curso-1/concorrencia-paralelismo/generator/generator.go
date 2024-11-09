package main

import (
	"fmt"

	"github.com/Wilson-cmd/html"
)

// // Canal somente leitura: <-chan

// func titulo(urls ...string) <-chan string {
// 	c := make(chan string)
// 	for _, url := range urls {
// 		go func(url string) {
// 			response, _ := http.Get(url)
// 			html, _ := io.ReadAll(response.Body)

// 			rgx := regexp.MustCompile(`<title>(.*?)</title`)

// 			c <- rgx.FindStringSubmatch(string(html))[0]
// 		}(url)
// 	}
// 	return c

// }

func main() {

	t1 := html.Titulo("https://google.com", "https://youtube.com")
	// t2 := titulo("https://cod3r.com.br", "https://amazon.com") //Não tem title
	fmt.Println(<-t1)
}
