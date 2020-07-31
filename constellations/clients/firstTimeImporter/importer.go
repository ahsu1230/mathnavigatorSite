package main

import (
	"flag"
	"log"
	"strings"
)

DEFAULT_LOCAL_ORION_HOST := "http://localhost:8001"

// This CLI reads JSON files (given a single file or a directory/folder) and 
// sends appropriate http requests to a specified orion webserver. If no orion
// is specified, it will use the localhost orion webserver.
// Make sure orion is healthy before running this CLI.
// Note: this importer does NOT override data. It only adds non-existing data.
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

	isFolder := folderPath != ""
	isFile := filePath != ""

	if (!isFolder && !isFile) {
		// error - must specify at least one. Exit program.
	}

	if (isFolder && isFile) {
		// error - cannot specify both. Exit program.
	}

	if (isFile) {

		if (strings.contains(filePath, "programs.json")) {
			importer.ImportProgram(orionHost, filePath)
		}
		// Determine what domain to use based on file name
		// Need a map between fileName -> domain type?
		// Read JSON file
		// Unmarshal JSON as a list of Domains
		// Need a map between domain type -> API url (?)
		// Send Http Request per domain
	} else { // is folder
		// for every file in folder, call filePath logic
		// if file does not have a correct associated domain, skip file
		// and print / log a warning that you're skipping the file
	}
}

func ReadJsonFile(jsonFilePath string) interface{}, error {
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	return byteValue, err
}