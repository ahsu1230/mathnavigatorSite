package main

import (
	"flag"
	"log"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

DEFAULT_LOCAL_ORION_HOST := "http://localhost:8001"

// This CLI reads JSON files (given a single file or a directory/folder) and 
// sends appropriate http requests to a specified orion webserver. If no orion
// is specified, it will use the localhost orion webserver.
// Make sure orion is healthy before running this CLI.
//
// You can run this CLI using:
// go run importer.go --folder jsons
// go run importer.go --folder jsons --orionHost https://www.andymathnavigator.com
// go run importer.go --file jsons/programs.json --domain program
// go run importer.go --file jsons/programs.json --domain program --orionHost https://www.andymathnavigator.com
func main() {
	log.Println("Orion First-Time-Importer Client starting...")

	orionHost := flag.String("orionHost", DEFAULT_LOCAL_ORION_HOST, "Host address of an Orion webserver")
	folderPath := flag.String("folder", "", "Path to folder that contains valid domain JSON files")
	filePath := flag.String("file", "", "Path to file that contains valid domain JSON");

	if (folderPath && filePath) {
		// error - cannot specify both. Exit program.
	}

	if (!folderPath && !filePath) {
		// error - must specify at least one. Exit program.
	}

	if (filePath) {
		// Determine what domain to use based on file name
		// Need a map between fileName -> domain type?
		// Read JSON file
		// Unmarshal JSON as a list of Domains
		// Need a map between domain type -> API url
		// Send Http Request per domain
		
		// Do we need to space out queries?
		// Should we first check which domains we don't need to add?
		// (This importer does NOT override data - it only adds non-existing data)
		// Do we need bulk creation endpoints on orion?
	}

	if (folderPath) {
		// for every file in folder, call filePath logic
		// if file does not have a correct associated domain, skip file
		// and print / log a warning that you're skipping the file
	}
}

func sendPostRequest(url string, body io.Reader) {
	resp, err := http.Post(url, "application/json; charset=UTF-8", body)
	if err != nil {
		log.Println("Post request was not fulfilled.", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Println("Response status was not successful.", resp)
	}
}