package integration_tests

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func resetSessionTables(t *testing.T) {
	resetTable(t, domains.TABLE_SESSIONS)
	resetTable(t, domains.TABLE_CLASSES)
	resetTable(t, domains.TABLE_SEMESTERS)
	resetTable(t, domains.TABLE_LOCATIONS)
	resetTable(t, domains.TABLE_PROGRAMS)
}

// Test: Create 3 Sessions, 2 With Same Class Id, and GetAllByClassId()
func Test_CreateSessions(t *testing.T) {
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
	body1 := createJsonBody(prog1)
	body2 := createJsonBody(prog2)
	body3 := createJsonBody(loc1)
	body4 := createJsonBody(semester1)
	body5 := createJsonBody(semester2)
	body6 := createJsonBody(&class1)
	body7 := createJsonBody(&class2)
	body8 := createJsonBody(session1)
	body9 := createJsonBody(session2)
	body10 := createJsonBody(session3)
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

	// Call Get All!
	recorder11 := sendHttpRequest(t, http.MethodGet, "/api/sessions/v1/class/fast_track_2020_spring_class_A", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder11.Code)
	var sessions []domains.Session
	if err := json.Unmarshal(recorder11.Body.Bytes(), &sessions); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, sessions[0].Id)
	assert.EqualValues(t, "fast_track_2020_spring_class_A", sessions[0].ClassId)
	assert.EqualValues(t, 1, sessions[1].Id)
	assert.EqualValues(t, "fast_track_2020_spring_class_A", sessions[1].ClassId)
	assert.EqualValues(t, 2, len(sessions))
}

// Test: Create 1 Session, Update it, GetBySessionId()
func Test_UpdateSession(t *testing.T) {
	resetSessionTables(t)

	// Create 1 Session
	start := time.Now().UTC()
	end := start.Add(time.Hour)
	prog1 := createProgram("fast_track", "Fast Track", 1, 12, "descript1")
	loc1 := createLocation("loc_1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	semester1 := createSemester("2020_spring", "Spring 2020")
	class1 := createClassUtil("fast_track", "2020_spring", "class_A", "loc_1", "5 pm - 7 pm", start, end)
	session1 := createSession("fast_track_2020_spring_class_A", start, end, false, "special lecture from guest")
	body1 := createJsonBody(prog1)
	body2 := createJsonBody(loc1)
	body3 := createJsonBody(semester1)
	body4 := createJsonBody(&class1)
	body5 := createJsonBody(session1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body3)
	recorder4 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body4)
	recorder5 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body5)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Update
	updatedSession := createSession("fast_track_2020_spring_class_A", start, end, true, "cancelled due to corona")
	updatedBody := createJsonBody(updatedSession)
	recorder6 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/session/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder6.Code)

	// Get
	recorder7 := sendHttpRequest(t, http.MethodGet, "/api/sessions/v1/session/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder7.Code)

	// Validate results
	var session domains.Session
	if err := json.Unmarshal(recorder7.Body.Bytes(), &session); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, session.Id)
	assert.EqualValues(t, "fast_track_2020_spring_class_A", session.ClassId)
	assert.EqualValues(t, "cancelled due to corona", session.Notes)
}

// Test: Create 1 Session, Delete it, GetBySessionId()
func Test_DeleteSession(t *testing.T) {
	resetSessionTables(t)

	// Create
	start := time.Now().UTC()
	end := start.Add(time.Hour)
	prog1 := createProgram("fast_track", "Fast Track", 1, 12, "descript1")
	loc1 := createLocation("loc_1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	semester1 := createSemester("2020_spring", "Spring 2020")
	class1 := createClassUtil("fast_track", "2020_spring", "class_A", "loc_1", "5 pm - 7 pm", start, end)
	session1 := createSession("fast_track_2020_spring_class_A", start, end, false, "special lecture from guest")
	body1 := createJsonBody(prog1)
	body2 := createJsonBody(loc1)
	body3 := createJsonBody(semester1)
	body4 := createJsonBody(&class1)
	body5 := createJsonBody(session1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body3)
	recorder4 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body4)
	recorder5 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body5)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Delete
	recorder6 := sendHttpRequest(t, http.MethodDelete, "/api/sessions/v1/session/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder6.Code)

	// Get
	recorder7 := sendHttpRequest(t, http.MethodGet, "/api/sessions/v1/session/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder7.Code)
}

// Helper methods
func createSession(classId string, startsAt time.Time, endsAt time.Time, canceled bool, notes string) domains.Session {
	return domains.Session{
		ClassId:  classId,
		StartsAt: startsAt,
		EndsAt:   endsAt,
		Canceled: canceled,
		Notes:    notes,
	}
}

func createClassUtil(programId string, semesterId string, classKey string, locationId string, times string, startDate time.Time, endDate time.Time) domains.Class {
	return domains.Class{
		ProgramId:  programId,
		SemesterId: semesterId,
		ClassKey:   domains.NewNullString(classKey),
		LocationId: locationId,
		Times:      times,
		StartDate:  startDate,
		EndDate:    endDate,
	}
}
