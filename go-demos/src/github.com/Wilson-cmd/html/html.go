package html

import (
	"io"
	"net/http"
	"regexp"
)

// Titulo obtem o titulo de uma pagina html
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			response, _ := http.Get(url)
			html, _ := io.ReadAll(response.Body)

			rgx := regexp.MustCompile(`<title>(.*?)</title`)

			c <- rgx.FindStringSubmatch(string(html))[0]
		}(url)
	}
	return c

}
