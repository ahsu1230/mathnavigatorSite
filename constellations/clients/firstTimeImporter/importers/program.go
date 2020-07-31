package importer

import (
	"encoding/json"
	"log"
)

// Struct that matches Program row/json from file
type ProgramFromFile struct {
	ProgramId   string `json:"programId"`
	Title       string `json:"title"`
	Grade1      uint   `json:"grade1"`
	Grade2      uint   `json:"grade2"`
	Description string `json:"description"`
}

// Need to copy from orion/domains
// Consider publishing Go modules to be able to be used here
type Program struct {
	ProgramId   string `json:"programId"`
	Name        string `json:"name"`
	Grade1      uint   `json:"grade1"`
	Grade2      uint   `json:"grade2"`
	Description string `json:"description"`
}

func ImportProgram(hostAddress string, data []byte) error {

	var programs []ProgramFromFile
	err := json.Unmarshal(data, &programs)
	if err != nil {
		log.Println("Error converting file to JSON")
		return err
	}

	for i := 0; i < len(programs); i++ {
		program := convert(programs[i])
		programJson, err := json.Marshal(program)
		if err != nil {
			log.Println("Error converting to JSON. Skipping", program.Name)
			continue
		}
		log.Println("Attempting to http POST request for", program.Name)
		SendPostRequest(hostAddress+"/api/programs/create", programJson)
	}

	return nil
}

func convert(oldProgram ProgramFromFile) Program {
	return Program{
		oldProgram.ProgramId,
		oldProgram.Title,
		oldProgram.Grade1,
		oldProgram.Grade2,
		oldProgram.Description,
	}
}
