package main

import (
	"encoding/base64"
	"net/http"
	"os"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// client := &http.Client{}
	// PATH := "reader-pdf-file.jpg"
	// PATH := "DIPLOMA-FRENTE-01-pb.jpg"
	PATH := "image.png"
	// uniformResourceLocator := "https://document-ai-extract-information.eastus2.inference.ml.azure.com/score"
	// bearer := "Bearer H43Z4vO3v0qp9dMouu5QR7JcSDpMjRaC"
	var base64Encoding string

	byts, err := os.ReadFile(PATH)
	checkError(err)

	mimeType := http.DetectContentType(byts)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(byts)

	os.WriteFile(".txt", []byte(base64Encoding), 0644)

	// data := []byte(base64Encoding)
	// request, err := http.NewRequest("POST", uniformResourceLocator, bytes.NewBuffer(data))
	// checkError(err)

	// request.Header.Add("Content-Type", "text/plain")
	// request.Header.Add("Authorization", bearer)
	// response, err := client.Do(request)
	// checkError(err)

	// defer response.Body.Close()
	// body, err := io.ReadAll(response.Body)
	// checkError(err)
	// fmt.Printf("%q\n", string(body))

}
