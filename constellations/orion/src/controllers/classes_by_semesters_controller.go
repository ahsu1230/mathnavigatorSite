package controllers

import (
	"errors"
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllProgramsSemestersClasses(c *gin.Context) {
	var listResults []domains.ProgramClassesBySemester

	// Fetch programs, semesters, classes from repo functions
	publishedOnly := ParseParamPublishedOnly(c)

	programs, err := repos.ProgramRepo.SelectAll()
	semesters, err := repos.SemesterRepo.SelectAll()
	classes, err := repos.ClassRepo.SelectAll(publishedOnly)

	// Convert lists into maps
	programMap := createProgramMap(programs)

	// Loop over semesterIds
	for i := 0; i < len(semesters); i++ {
		// Create Semester Object
		semesterObj := semesters[i]

		// Make ProgramClass structs
		programClasses, err := createProgramClassesForSemester(semesters[i].SemesterId, programs, classes, programMap)
		if err != nil {
			c.Error(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Make ProgramClassesBySemester struct
		programClassesBySemester := domains.ProgramClassesBySemester{
			Semester:       semesterObj,
			ProgramClasses: programClasses,
		}

		listResults = append(listResults, programClassesBySemester)
	}

	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &listResults)
	}
}

// Helper functions
func createProgramClassesForSemester(semesterId string, programs []domains.Program, classes []domains.Class, programMap map[string]domains.Program) ([]domains.ProgramClass, error) {
	var programClasses []domains.ProgramClass

	// Create a list of all classes in one semester classSlice
	classSlice := make([]domains.Class, 0)
	for i := 0; i < len(classes); i++ {
		if classes[i].SemesterId == semesterId {
			classSlice = append(classSlice, classes[i])
		}
	}

	// Create map mapping programId to list of classes programClassMap
	programClassMap, err := createProgramClassMap(classSlice, programMap)
	if err != nil {
		return []domains.ProgramClass{}, err
	}

	// Create list of ProgramClass
	for i := 0; i < len(programs); i++ {
		programId := programs[i].ProgramId
		programClass := domains.ProgramClass{
			ProgramObj: programs[i],
			Classes:    programClassMap[programId],
		}
		programClasses = append(programClasses, programClass)
	}

	return programClasses, nil
}

func createProgramClassMap(classSlice []domains.Class, programMap map[string]domains.Program) (map[string][]domains.Class, error) {
	// Create map mapping programId to list of classes programClassMap
	programClassMap := make(map[string][]domains.Class)

	// For each class, get programId and put into programClassMap
	for i := 0; i < len(classSlice); i++ {
		programId := classSlice[i].ProgramId
		if _, ok := programMap[programId]; !ok {
			err := errors.New("programId not found in list of programs")
			return map[string][]domains.Class{}, err
		}

		if _, ok := programClassMap[programId]; ok {
			programClassMap[programId] = append(programClassMap[programId], classSlice[i])
		} else {
			programClassMap[programId] = []domains.Class{classSlice[i]}
		}
	}
	return programClassMap, nil
}

func Find(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}

func createProgramMap(programs []domains.Program) map[string]domains.Program {
	programMap := make(map[string]domains.Program)
	for i := 0; i < len(programs); i++ {
		programMap[programs[i].ProgramId] = programs[i]
	}
	return programMap
}
