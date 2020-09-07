package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var HOST_ADDRESS string

type IdResp struct {
	Id uint `json:"id"`
}

type IdsResp struct {
	Ids []uint `json:"ids"`
}

func InitHostAddress(hostAddress string) {
	HOST_ADDRESS = hostAddress
}

func SendPostRequest(urlPath string, body io.Reader) []byte {
	fullUrl := HOST_ADDRESS + urlPath
	resp, err := http.Post(fullUrl, "application/json; charset=UTF-8", body)
	if err != nil {
		log.Println("Post request was not fulfilled.", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Println("Response status was not successful.", resp)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	return bytes
}

// Assume the response body looks like:
// { "id": ___ }
func GetIdFromBody(bytes []byte) (uint, error) {
	log.Printf("*** %s\n", string(bytes))
	var resp IdResp
	if err := json.Unmarshal(bytes, &resp); err != nil {
		log.Printf("unexpected error: %v\n", err)
		return 0, err
	}
	return resp.Id, nil
}

func GetIdsFromBody(bytes []byte) ([]uint, error) {
	log.Printf("*** %s\n", string(bytes))
	var resp IdsResp
	if err := json.Unmarshal(bytes, &resp); err != nil {
		log.Printf("unexpected error: %v\n", err)
		return []uint{}, err
	}
	return resp.Ids, nil
}
