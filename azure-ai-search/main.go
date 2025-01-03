package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// Documentation
/*https://learn.microsoft.com/pt-br/rest/api/searchservice/*/
var client = &http.Client{}

func main() {

	URL := "https://<name>.search.windows.net/indexes/<name-index>/docs/search?api-version=2024-11-01-preview"
	body, _ := json.Marshal(map[string]string{

		"search": "O que diz no memorando mais recente, retorne 3 linhas",
		"count":  "true",
		// "queryType": "semantic",
		// "semanticConfiguration": "my-semantic-config",
		// "captions":              "extractive",
		// "answers":               "extractive|count-3",
		// "queryLanguage":         "en-us",
	})

	payload := bytes.NewBuffer(body)

	if request, e := http.NewRequest(http.MethodPost,
		URL, payload); e != nil {
		log.Println(e.Error())
	} else {
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("api-key", "<key>")

		if response, err := client.Do(request); err != nil {
			log.Println(err.Error())
		} else {
			b, er := io.ReadAll(response.Body)

			if er != nil {
				log.Println(er.Error())
				return
			}

			defer response.Body.Close()

			os.WriteFile("response.json", b, 0644)

		}

	}
}
