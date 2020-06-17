package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// This CLI sends http requests to a local orion webserver.
// Make sure orion is healthy before running this CLI
//
// You can run this CLI using:
// go run filler.go
//
// or via a binary:
// go build filler.go
// ./filler

func main() {
	fmt.Println("Orion Fake Data Filler Client starting...")

	runFiller("http://localhost:8001")

	fmt.Println("Done filling orion")
}

func runFiller(hostAddress string) {
	body := strings.NewReader(`{
		"programId": "ap_calculus",
		"name": "AP Calculus",
		"grade1": 9,
		"grade2": 12,
		"description": "Students should take this course if they aim to take the AP Calculus Exam"
	}`)
	sendPostRequest(hostAddress+"/api/programs/create", body)
}

func sendPostRequest(url string, body io.Reader) {
	resp, err := http.Post(url, "application/json; charset=UTF-8", body)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Fatalln("Response status was not successful.", resp)
	}
}
