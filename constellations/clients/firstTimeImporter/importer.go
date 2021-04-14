package main

import (
	"./importers"
	"flag"
	"io/ioutil"
	"log"
	"strings"
)

var DEFAULT_LOCAL_ORION_HOST string = "http://localhost:8001"

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

	var orionHost, folderPath, filePath string
	flag.StringVar(&orionHost, "orionHost", DEFAULT_LOCAL_ORION_HOST, "Host address of an Orion webserver")
	flag.StringVar(&folderPath, "folder", "", "Path to folder that contains valid domain JSON files")
	flag.StringVar(&filePath, "file", "", "Path to file that contains valid domain JSON")
	flag.Parse()

	isFolder := (folderPath != "")
	isFile := (filePath != "")

	if !isFolder && !isFile {
		// error - must specify at least one. Exit program.
	}

	if isFolder && isFile {
		// error - cannot specify both. Exit program.
	}

	if isFile {
		data, err := readJsonFile(filePath)
		if err != nil {
			// error reading json file
			return
		}
		log.Println(data)

		if strings.Contains(filePath, "programs.json") {
			err = importer.ImportProgram(orionHost, data)
			if err != nil {
				// error importing program
				return
			}
		}
	} else { // is folder
		// for every file in folder, call filePath logic
		// if file does not have a correct associated domain, skip file
		// and print / log a warning that you're skipping the file
	}
}

func readJsonFile(jsonFilePath string) ([]byte, error) {
	byteValue, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return nil, err
	}
	return byteValue, err
}
