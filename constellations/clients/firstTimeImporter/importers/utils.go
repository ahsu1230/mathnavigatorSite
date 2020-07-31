package importer

import (
	"bytes"
	"log"
	"net/http"
)

func SendPostRequest(url string, body []byte) {
	resp, err := http.Post(url, "application/json; charset=UTF-8", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Post request was not fulfilled.", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Println("Response status was not successful.", resp.StatusCode)
		// print response body
	}
}
