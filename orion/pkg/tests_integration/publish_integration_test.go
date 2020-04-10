package integration_tests

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

// Test: GetUnpublished()
func Test_GetUnpublishedAchievements(t *testing.T) {
	resetTable(t, domains.TABLE_ACHIEVEMENTS)

	achieve1 := createAchievement(2020, "message1")
	achieve2 := createAchievement(2021, "message2")
	body1 := createJsonBody(&achieve1)
	body2 := createJsonBody(&achieve2)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Call Get Unpublished!
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder3.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, unpublishedDomains.Achieves[0].Id)
	assert.EqualValues(t, 2020, unpublishedDomains.Achieves[0].Year)
	assert.EqualValues(t, "message1", unpublishedDomains.Achieves[0].Message)
	assert.EqualValues(t, 2, unpublishedDomains.Achieves[1].Id)
	assert.EqualValues(t, 2021, unpublishedDomains.Achieves[1].Year)
	assert.EqualValues(t, "message2", unpublishedDomains.Achieves[1].Message)
	assert.EqualValues(t, 2, len(unpublishedDomains.Achieves))
}

// Test: Get All Unpublished Programs, Publish a Few, then Get All Unpublished Again
func Test_PublishPrograms(t *testing.T) {
	resetTable(t, domains.TABLE_PROGRAMS)

	// Create
	program1 := createProgram("prog1", "Program1", 2, 3, "descript1")
	program2 := createProgram("prog2", "Program2", 8, 12, "descript2")
	program3 := createProgram("prog3", "Program3", 1, 12, "descript3")
	body1 := createJsonBody(&program1)
	body2 := createJsonBody(&program2)
	body3 := createJsonBody(&program3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Get All Unpublished
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)

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
	publishBody := createJsonBody(publishIds)
	recorder5 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/publish", publishBody)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Get All Unpublished Again
	recorder6 := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder6.Code)
	if err := json.Unmarshal(recorder6.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "prog2", unpublishedDomains.Programs[0].ProgramId)
	assert.EqualValues(t, 1, len(unpublishedDomains.Programs))
}

// Test: Get All Unpublished Locations, Publish a Few, then Get All Unpublished Again
func Test_PublishLocations(t *testing.T) {
	resetTable(t, domains.TABLE_LOCATIONS)

	// Create
	location1 := createLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	location2 := createLocation("loc2", "4040 Location Ave", "Dity", "MD", "77294-1243", "Room 2")
	location3 := createLocation("loc3", "4040 Location Blvd", "Eity", "ND", "08430-0302", "Room 3")
	body1 := createJsonBody(&location1)
	body2 := createJsonBody(&location2)
	body3 := createJsonBody(&location3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Get All Unpublished
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder4.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", unpublishedDomains.Locations[0].LocId)
	assert.EqualValues(t, "loc2", unpublishedDomains.Locations[1].LocId)
	assert.EqualValues(t, "loc3", unpublishedDomains.Locations[2].LocId)
	assert.EqualValues(t, 3, len(unpublishedDomains.Locations))

	// Publish loc1 and loc3
	publishIds := []string{"loc1", "loc3"}
	publishBody := createJsonBody(publishIds)
	recorder5 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/publish", publishBody)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Get All Unpublished Again
	recorder6 := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder6.Code)
	if err := json.Unmarshal(recorder6.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc2", unpublishedDomains.Locations[0].LocId)
	assert.EqualValues(t, 1, len(unpublishedDomains.Locations))
}

// Test: Get All Unpublished Sessions, Publish a Few, then Get All Unpublished Again
func Test_PublishSessions(t *testing.T) {
	resetSessionTables(t)

	// Create
	start := time.Now().UTC()
	mid := start.Add(time.Minute * 30)
	end := start.Add(time.Hour)
	prog1 := createProgram("fast_track", "Fast Track", 1, 12, "descript1")
	prog2 := createProgram("slow_track", "Slow Track", 1, 12, "descript1")
	loc1 := createLocation("loc_1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	semester1 := createSemester("2020_spring", "Spring 2020")
	semester2 := createSemester("2020_fall", "Fall 2020")
	class1 := createClassUtil("fast_track", "2020_spring", "class_A", "loc_1", "5 pm - 7 pm", start, end)
	class2 := createClassUtil("slow_track", "2020_fall", "class_B", "loc_1", "3 pm - 7 pm", start, end)
	session1 := createSession("fast_track_2020_spring_class_A", mid, end, false, "special lecture from guest")
	session2 := createSession("fast_track_2020_spring_class_A", start, end, true, "May 5th regular meeting")
	session3 := createSession("slow_track_2020_fall_class_B", start, end, false, "May 5th regular meeting")
	body1 := createJsonBody(&prog1)
	body2 := createJsonBody(&prog2)
	body3 := createJsonBody(&loc1)
	body4 := createJsonBody(semester1)
	body5 := createJsonBody(semester2)
	body6 := createJsonBody(&class1)
	body7 := createJsonBody(&class2)
	body8 := createJsonBody(&session1)
	body9 := createJsonBody(&session2)
	body10 := createJsonBody(&session3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body3)
	recorder4 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body4)
	recorder5 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body5)
	recorder6 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body6)
	recorder7 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body7)
	recorder8 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body8)
	recorder9 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body9)
	recorder10 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body10)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)
	assert.EqualValues(t, http.StatusOK, recorder6.Code)
	assert.EqualValues(t, http.StatusOK, recorder7.Code)
	assert.EqualValues(t, http.StatusOK, recorder8.Code)
	assert.EqualValues(t, http.StatusOK, recorder9.Code)
	assert.EqualValues(t, http.StatusOK, recorder10.Code)

	// Get All Unpublished
	recorder11 := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder11.Code)
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder11.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, unpublishedDomains.Sessions[0].Id)
	assert.EqualValues(t, 2, unpublishedDomains.Sessions[1].Id)
	assert.EqualValues(t, 3, unpublishedDomains.Sessions[2].Id)
	assert.EqualValues(t, 3, len(unpublishedDomains.Sessions))

	// Publish first and third session
	publishIds := []uint{1, 3}
	publishBody := createJsonBody(publishIds)
	recorder12 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/publish", publishBody)
	assert.EqualValues(t, http.StatusOK, recorder12.Code)

	// Get All Unpublished Again
	recorder13 := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder13.Code)
	if err := json.Unmarshal(recorder13.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, unpublishedDomains.Sessions[0].Id)
	assert.EqualValues(t, 1, len(unpublishedDomains.Sessions))
}
