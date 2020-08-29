package utils

import (
	"io"
	"log"
	"net/http"
)

func SendPostRequest(url string, body io.Reader) {
	resp, err := http.Post(url, "application/json; charset=UTF-8", body)
	if err != nil {
		log.Println("Post request was not fulfilled.", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Println("Response status was not successful.", resp)
	}
}
