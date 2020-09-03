package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/cache"
	"github.com/gin-gonic/gin"
)

func GetAllProgramsSemestersClasses(c *gin.Context) {
	utils.LogControllerMethod(c, "controller.GetAllProgramsSemestersClasses")
	var listResults []domains.ProgramClassesBySemester

	// Fetch programs, semesters, classes from repo functions
	publishedOnly := utils.ParseParamPublishedOnly(c)

	// Fetch from cache first to save on computation work
	cachedList, err := cache.GetAllProgramClassesBySemester()
	if err == nil {
		c.JSON(http.StatusOK, &cachedList)
		return
	}
	// Ignore err. Cache miss means we compute normally

	programs, err := repos.ProgramRepo.SelectAll()
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}

	semesters, err := repos.SemesterRepo.SelectAll()
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}

	classes, err := repos.ClassRepo.SelectAll(publishedOnly)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}

	// Convert lists into maps
	programMap := createProgramMap(programs)

	// Loop over semesterIds
	for i := 0; i < len(semesters); i++ {
		// Create Semester Object
		semesterObj := semesters[i]

		// Make ProgramClass structs
		programClasses, err := createProgramClassesForSemester(semesters[i].SemesterId, programs, classes, programMap)
		if err != nil {
			c.Error(appErrors.WrapRepo(err))
			return
		}

		// Make ProgramClassesBySemester struct
		programClassesBySemester := domains.ProgramClassesBySemester{
			Semester:       semesterObj,
			ProgramClasses: programClasses,
		}

		listResults = append(listResults, programClassesBySemester)
	}
	cache.SetAllProgramClassesBySemester(listResults)
	c.JSON(http.StatusOK, &listResults)
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
		if _, ok := programClassMap[programId]; ok {
			programClass := domains.ProgramClass{
				ProgramObj: programs[i],
				Classes:    programClassMap[programId],
			}
			programClasses = append(programClasses, programClass)
		}
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
			err := appErrors.WrapCtrlf("programId %s not found in list of programs", programId)
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

func createProgramMap(programs []domains.Program) map[string]domains.Program {
	programMap := make(map[string]domains.Program)
	for i := 0; i < len(programs); i++ {
		programMap[programs[i].ProgramId] = programs[i]
	}
	return programMap
}
