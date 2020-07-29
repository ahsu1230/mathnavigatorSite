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

	// Fetch progrmas, semesters, classes from repo functions
	publishedOnly := ParseParamPublishedOnly(c)

	programs, err := repos.ProgramRepo.SelectAll(publishedOnly)
	semesters, err := repos.SemesterRepo.SelectAll(publishedOnly)
	classes, err := repos.ClassRepo.SelectAll(publishedOnly)

	// Convert lists into maps
	programMap := createProgramMap(programs)

	// Loop over semesterIds
	for i := 0; i < len(semesters); i++ {
		// Create Semester Object
		semesterObj := semesters[i]

		programClasses := createProgramClassesForSemester(c, semesters[i].SemesterId, programs, classes, programMap)

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

func createProgramClassesForSemester(c *gin.Context, semesterId string, programs []domains.Program, classes []domains.Class, programMap map[string]domains.Program) []domains.ProgramClass {
	var programClasses []domains.ProgramClass
	// Create a list of all classes in one semester classSlice
	classSlice := make([]domains.Class, 0)
	for i := 0; i < len(classes); i++ {
		if classes[i].SemesterId == semesterId {
			classSlice = append(classSlice, classes[i])
		}
	}

	// Create map mapping programId to list of classes programClassMap
	programClassMap := make(map[string][]domains.Class)

	// For each class, get programId and put into programClassMap
	for i := 0; i < len(classSlice); i++ {
		programId := classSlice[i].ProgramId
		if _, ok := programMap[programId]; !ok {
			err := errors.New("programId not found in list of programs")
			c.Error(err)
			c.String(http.StatusInternalServerError, err.Error())
		}

		if _, ok := programClassMap[programId]; ok {
			programClassMap[programId] = append(programClassMap[programId], classSlice[i])
		} else {
			programClassMap[programId] = []domains.Class{classSlice[i]}
		}
	}

	// Account for programs that are in map but have no classes
	for key, _ := range programMap {
		if _, ok := programClassMap[key]; ok {
			continue
		} else {
			programClassMap[key] = []domains.Class{}
		}
	}

	// Create list of ProgramClass
	for i := 0; i < len(programs); i++ {
		programId := programs[i].ProgramId
		programClasses = append(programClasses, updateProgramClass(programs[i], programClassMap[programId]))
	}

	return programClasses
}

func updateProgramClass(programObj domains.Program, classes []domains.Class) domains.ProgramClass {
	return domains.ProgramClass{
		ProgramObj: programObj,
		Classes:    classes,
	}
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
