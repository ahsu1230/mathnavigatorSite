package tests_integration

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Get All Unpublished Programs, Publish a Few, then Get All Unpublished Again
func Test_PublishPrograms(t *testing.T) {
	// Create
	program1 := createProgram("prog1", "Program1", 2, 3, "descript1", 0)
	program2 := createProgram("prog2", "Program2", 8, 12, "descript2", 1)
	program3 := createProgram("prog3", "Program3", 1, 12, "descript3", 0)
	body1 := utils.CreateJsonBody(&program1)
	body2 := utils.CreateJsonBody(&program2)
	body3 := utils.CreateJsonBody(&program3)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Get All Unpublished
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder4.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "prog1", unpublishedDomains.Programs[0].ProgramId)
	assert.EqualValues(t, "prog2", unpublishedDomains.Programs[1].ProgramId)
	assert.EqualValues(t, "prog3", unpublishedDomains.Programs[2].ProgramId)
	assert.EqualValues(t, 3, len(unpublishedDomains.Programs))

	// Publish prog1 and prog3
	publishIds := []string{"prog1", "prog3"}
	publishBody := utils.CreateJsonBody(publishIds)
	recorder5 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/publish", publishBody)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Get All Unpublished Again
	recorder6 := utils.SendHttpRequest(t, http.MethodGet, "/api/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder6.Code)
	if err := json.Unmarshal(recorder6.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "prog2", unpublishedDomains.Programs[0].ProgramId)
	assert.EqualValues(t, 1, len(unpublishedDomains.Programs))

	// Get Published Only
	recorder7 := utils.SendHttpRequest(t, http.MethodGet, "/api/programs/all?published=true", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder7.Code)
	var programs []domains.Program
	if err := json.Unmarshal(recorder7.Body.Bytes(), &programs); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "prog1", programs[0].ProgramId)
	assert.EqualValues(t, "prog3", programs[1].ProgramId)
	assert.EqualValues(t, 2, len(programs))

	utils.ResetTable(t, domains.TABLE_PROGRAMS)
}

// Test: Create 2 Classes and Publish 1
func Test_PublishClasses(t *testing.T) {
	// Create
	createAllProgramsSemestersLocations(t)
	class1 := createClass(1)
	class2 := createClass(2)
	body1 := utils.CreateJsonBody(&class1)
	body2 := utils.CreateJsonBody(&class2)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get All Published
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var classes1 []domains.Class
	if err := json.Unmarshal(recorder3.Body.Bytes(), &classes1); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 0, len(classes1))

	// Publish
	classIds := []string{"program1_2020_spring_class2"}
	body3 := utils.CreateJsonBody(&classIds)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/publish", body3)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Get All Published
	recorder5 := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Validate results
	var classes2 []domains.Class
	if err := json.Unmarshal(recorder5.Body.Bytes(), &classes2); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertClass(t, 2, classes2[0])
	assert.EqualValues(t, 1, len(classes2))

	// Get All Unpublished
	recorder6 := utils.SendHttpRequest(t, http.MethodGet, "/api/unpublished", nil)
	assert.EqualValues(t, http.StatusOK, recorder6.Code)

	// Validate results
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder6.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertClass(t, 1, unpublishedDomains.Classes[0])
	assert.EqualValues(t, 1, len(unpublishedDomains.Classes))

	resetClassTables(t)
}

// Test: Get All Unpublished Locations, Publish a Few, then Get All Unpublished Again
func Test_PublishLocations(t *testing.T) {
	// Create
	location1 := createLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	location2 := createLocation("loc2", "4040 Location Ave", "Dity", "MD", "77294-1243", "Room 2")
	location3 := createLocation("loc3", "4040 Location Blvd", "Eity", "ND", "08430-0302", "Room 3")
	body1 := utils.CreateJsonBody(&location1)
	body2 := utils.CreateJsonBody(&location2)
	body3 := utils.CreateJsonBody(&location3)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Get All Unpublished
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder4.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", unpublishedDomains.Locations[0].LocationId)
	assert.EqualValues(t, "loc2", unpublishedDomains.Locations[1].LocationId)
	assert.EqualValues(t, "loc3", unpublishedDomains.Locations[2].LocationId)
	assert.EqualValues(t, 3, len(unpublishedDomains.Locations))

	// Publish loc1 and loc3
	publishIds := []string{"loc1", "loc3"}
	publishBody := utils.CreateJsonBody(publishIds)
	recorder5 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/publish", publishBody)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Get All Unpublished Again
	recorder6 := utils.SendHttpRequest(t, http.MethodGet, "/api/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder6.Code)
	if err := json.Unmarshal(recorder6.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc2", unpublishedDomains.Locations[0].LocationId)
	assert.EqualValues(t, 1, len(unpublishedDomains.Locations))

	// Get Published Only
	recorder7 := utils.SendHttpRequest(t, http.MethodGet, "/api/locations/all?published=true", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder7.Code)
	var locations []domains.Location
	if err := json.Unmarshal(recorder7.Body.Bytes(), &locations); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", locations[0].LocationId)
	assert.EqualValues(t, "loc3", locations[1].LocationId)
	assert.EqualValues(t, 2, len(locations))

	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Test: Create 2 Achievements and Publish 1
func Test_PublishAchievements(t *testing.T) {
	// Create
	achieve1 := createAchievement(2020, "message1", 1)
	achieve2 := createAchievement(2021, "message2", 2)
	body1 := utils.CreateJsonBody(&achieve1)
	body2 := utils.CreateJsonBody(&achieve2)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get All Published
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/achievements/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var achieves1 []domains.Achieve
	if err := json.Unmarshal(recorder3.Body.Bytes(), &achieves1); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 0, len(achieves1))

	// Publish
	ids := []uint{1}
	body3 := utils.CreateJsonBody(&ids)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/publish", body3)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Get All Published
	recorder5 := utils.SendHttpRequest(t, http.MethodGet, "/api/achievements/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Validate results
	var achieves2 []domains.Achieve
	if err := json.Unmarshal(recorder5.Body.Bytes(), &achieves2); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieves2[0].Id)
	assert.EqualValues(t, 2020, achieves2[0].Year)
	assert.EqualValues(t, "message1", achieves2[0].Message)
	assert.EqualValues(t, 1, len(achieves2))

	// Get All Unpublished
	recorder6 := utils.SendHttpRequest(t, http.MethodGet, "/api/unpublished", nil)
	assert.EqualValues(t, http.StatusOK, recorder6.Code)

	// Validate results
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder6.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, unpublishedDomains.Achieves[0].Id)
	assert.EqualValues(t, 2021, unpublishedDomains.Achieves[0].Year)
	assert.EqualValues(t, "message2", unpublishedDomains.Achieves[0].Message)
	assert.EqualValues(t, 1, len(unpublishedDomains.Achieves))

	utils.ResetTable(t, domains.TABLE_ACHIEVEMENTS)
}

// Test: Create 2 Semesters and Publish 1
func Test_PublishSemesters(t *testing.T) {
	// Create
	semester1 := createSemester("2020_fall", "Fall 2020", 1)
	semester2 := createSemester("2020_winter", "Winter 2020", 2)
	body1 := utils.CreateJsonBody(&semester1)
	body2 := utils.CreateJsonBody(&semester2)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get All Published
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/semesters/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var semesters1 []domains.Semester
	if err := json.Unmarshal(recorder3.Body.Bytes(), &semesters1); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 0, len(semesters1))

	// Publish
	semesterIds := []string{"2020_fall"}
	body3 := utils.CreateJsonBody(&semesterIds)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/publish", body3)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Get All Published
	recorder5 := utils.SendHttpRequest(t, http.MethodGet, "/api/semesters/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Validate results
	var semesters2 []domains.Semester
	if err := json.Unmarshal(recorder5.Body.Bytes(), &semesters2); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, semesters2[0].Id)
	assert.EqualValues(t, "2020_fall", semesters2[0].SemesterId)
	assert.EqualValues(t, "Fall 2020", semesters2[0].Title)
	assert.EqualValues(t, 1, len(semesters2))

	// Get All Unpublished
	recorder6 := utils.SendHttpRequest(t, http.MethodGet, "/api/unpublished", nil)
	assert.EqualValues(t, http.StatusOK, recorder6.Code)

	// Validate results
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder6.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, unpublishedDomains.Semesters[0].Id)
	assert.EqualValues(t, "2020_winter", unpublishedDomains.Semesters[0].SemesterId)
	assert.EqualValues(t, "Winter 2020", unpublishedDomains.Semesters[0].Title)
	assert.EqualValues(t, 1, len(unpublishedDomains.Semesters))

	utils.ResetTable(t, domains.TABLE_SEMESTERS)
}

// Test: Get All Unpublished Sessions, Publish a Few, then Get All Unpublished Again
func Test_PublishSessions(t *testing.T) {
	// Create
	start := time.Now().UTC()
	mid := start.Add(time.Minute * 30)
	end := start.Add(time.Hour)
	prog1 := createProgram("fast_track", "Fast Track", 1, 12, "descript1", 0)
	prog2 := createProgram("slow_track", "Slow Track", 1, 12, "descript1", 1)
	loc1 := createLocation("loc_1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	semester1 := createSemester("2020_spring", "Spring 2020", 1)
	semester2 := createSemester("2020_fall", "Fall 2020", 2)
	class1 := createClassUtil("fast_track", "2020_spring", "class_A", "loc_1", "5 pm - 7 pm", start, end)
	class2 := createClassUtil("slow_track", "2020_fall", "class_B", "loc_1", "3 pm - 7 pm", start, end)
	session1 := createSession("fast_track_2020_spring_class_A", mid, end, false, "special lecture from guest")
	session2 := createSession("fast_track_2020_spring_class_A", start, end, true, "May 5th regular meeting")
	session3 := createSession("slow_track_2020_fall_class_B", start, end, false, "May 5th regular meeting")
	body1 := utils.CreateJsonBody(&prog1)
	body2 := utils.CreateJsonBody(&prog2)
	body3 := utils.CreateJsonBody(&loc1)
	body4 := utils.CreateJsonBody(&semester1)
	body5 := utils.CreateJsonBody(&semester2)
	body6 := utils.CreateJsonBody(&class1)
	body7 := utils.CreateJsonBody(&class2)
	body8 := utils.CreateJsonBody([]domains.Session{session1, session2, session3})
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body3)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body4)
	recorder5 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body5)
	recorder6 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body6)
	recorder7 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body7)
	recorder8 := utils.SendHttpRequest(t, http.MethodPost, "/api/sessions/create", body8)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)
	assert.EqualValues(t, http.StatusOK, recorder6.Code)
	assert.EqualValues(t, http.StatusOK, recorder7.Code)
	assert.EqualValues(t, http.StatusOK, recorder8.Code)

	// Get All Unpublished
	recorder9 := utils.SendHttpRequest(t, http.MethodGet, "/api/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder9.Code)
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder9.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, unpublishedDomains.Sessions[0].Id)
	assert.EqualValues(t, 2, unpublishedDomains.Sessions[1].Id)
	assert.EqualValues(t, 3, unpublishedDomains.Sessions[2].Id)
	assert.EqualValues(t, 3, len(unpublishedDomains.Sessions))

	// Publish first and third session
	publishIds := []uint{1, 3}
	publishBody := utils.CreateJsonBody(publishIds)
	recorder10 := utils.SendHttpRequest(t, http.MethodPost, "/api/sessions/publish", publishBody)
	assert.EqualValues(t, http.StatusOK, recorder10.Code)

	// Get All Unpublished Again
	recorder11 := utils.SendHttpRequest(t, http.MethodGet, "/api/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder11.Code)
	if err := json.Unmarshal(recorder11.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, unpublishedDomains.Sessions[0].Id)
	assert.EqualValues(t, 1, len(unpublishedDomains.Sessions))

	// Get Published Only
	recorder12 := utils.SendHttpRequest(t, http.MethodGet, "/api/sessions/class/fast_track_2020_spring_class_A?published=true", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder12.Code)
	var sessions []domains.Session
	if err := json.Unmarshal(recorder12.Body.Bytes(), &sessions); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, sessions[0].Id)
	assert.EqualValues(t, 1, len(sessions))

	resetSessionTables(t)
}
