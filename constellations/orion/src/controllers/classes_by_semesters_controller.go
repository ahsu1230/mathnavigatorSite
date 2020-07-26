package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllProgramsSemestersClasses(c *gin.Context) {
	var classSemesterJson domains.ProgramClassesBySemester

	// Fetch progrmas, semesters, classes from repo functions
	publishedOnly := ParseParamPublishedOnly(c)

	programs, err := repos.ProgramRepo.SelectAll(publishedOnly)
	semesters, err := repos.SemesterRepo.SelectAll(publishedOnly)
	classes, err := repos.ClassRepo.SelectAll(publishedOnly)

	// Convert lists into maps
	programMap := make(map[string]domains.Program)
	semesterMap := make(map[string]domains.Semester)
	classMap := make(map[string]domains.Class)

	for i := 0; i < len(programs); i++ {
		programMap[programs[i].ProgramId] = programs[i]
	}
	for i := 0; i < len(semesters); i++ {
		semesterMap[semesters[i].SemesterId] = semesters[i]
	}
	for i := 0; i < len(classes); i++ {
		classMap[classes[i].ClassId] = classes[i]
	}

	// Semester to class list map (maps semesterId to Class object)
	semesterClassMap := make(map[string][]domains.Class)

	// Loop over semesterIds
	for _, value := range semesterMap {
		// Create Semester Object
		semesterObj := value
		var programClassStruct domains.ProgramClass
		var programObj domains.Program

		// Find classes that have the same semesterId and append to semesterClassMap
		classSlice := make([]domains.Class, 0)
		for i := 0; i < len(classes); i++ {
			if classes[i].SemesterId == semesterObj.SemesterId {
				classSlice = append(classSlice, classes[i])
			}
		}
		// List of classes in specific semester
		semesterClassMap[semesterObj.SemesterId] = classSlice

		// Create list of programIds in specific semester
		var programList []string
		for i := 0; i < len(classSlice); i++ {
			if Find(programList, classSlice[i].ProgramId) != -1 {
				continue
			} else {
				programList = append(programList, classSlice[i].ProgramId)
			}
		}

		// Find all classes in each program
		finalClassList := make([]domains.Class, 0)
		listProgramClass := make([]domains.ProgramClass, 0)
		for i := 0; i < len(programList); i++ {
			for j := 0; j < len(classSlice); j++ {
				if programList[i] == classSlice[j].ProgramId {
					finalClassList = append(finalClassList, classSlice[j])
				}
			}

			// Check if programId exists inside programMap
			if value, ok := programMap[programList[i]]; ok {
				programObj = value
			} else {
				c.Error(err)
				c.String(http.StatusInternalServerError, err.Error())
			}

			// Make ProgramClass struct
			programClassStruct = updateProgramClass(programObj, finalClassList)
			// Append to list of ProgramClasses
			listProgramClass = append(listProgramClass, programClassStruct)
			finalClassList = nil
		}

		// Make ProgramClassesBySemester struct
		classSemesterJson = updateProgramClassesBySemester(value, listProgramClass)
		c.BindJSON(&classSemesterJson)
	}

	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &classSemesterJson)
	}
	return
}

func updateProgramClass(programObj domains.Program, classes []domains.Class) domains.ProgramClass {
	return domains.ProgramClass{
		ProgramObj: programObj,
		Classes:    classes,
	}
}

func updateProgramClassesBySemester(semester domains.Semester, programClasses []domains.ProgramClass) domains.ProgramClassesBySemester {
	return domains.ProgramClassesBySemester{
		Semester:       semester,
		ProgramClasses: programClasses,
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
