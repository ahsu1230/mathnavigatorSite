package utils

import (
	"io"
	"log"
	"net/http"
)

var HOST_ADDRESS string

func InitHostAddress(hostAddress string) {
	HOST_ADDRESS = hostAddress
}

func SendPostRequest(urlPath string, body io.Reader) {
	fullUrl := HOST_ADDRESS + urlPath
	resp, err := http.Post(fullUrl, "application/json; charset=UTF-8", body)
	if err != nil {
		log.Println("Post request was not fulfilled.", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Println("Response status was not successful.", resp)
	}
}
