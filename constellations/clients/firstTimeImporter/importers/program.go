package importers

import (
	"log"
	"os"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func ImportProgram(hostAddress, jsonFilePath string) {

	var programs []domains.Program
	json.Unmarshal(byteValue, &programs)

	for i := 0; i < len(programs); i++ {
		fmt.Println(i, "::", programs[i])
		// send HTTP request to create a Program (found in utils)
	}

	return nil
}