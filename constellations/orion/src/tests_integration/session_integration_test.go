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

// Test: Create 3 Sessions, 2 With Same Class Id, and GetAllByClassId()
func TestCreateSessions(t *testing.T) {
	// Create
	start := time.Now().UTC()
	mid := start.Add(time.Minute * 30)
	end := start.Add(time.Hour)
	prog1 := createProgram("fast_track", "Fast Track", 1, 12, "descript1", 0)
	prog2 := createProgram("slow_track", "Slow Track", 1, 12, "descript1", 1)
	loc1 := createLocation("loc_1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	semester1 := createSemester("2020_spring", "Spring 2020", 1)
	semester2 := createSemester("2020_fall", "Fall 2020", 2)
	class1 := createClassUtil("fast_track", "2020_spring", "class_A", "loc_1", "5 pm - 7 pm", 50, 0)
	class2 := createClassUtil("slow_track", "2020_fall", "class_B", "loc_1", "3 pm - 7 pm", 50, 0)
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

	// Call Get All!
	recorder9 := utils.SendHttpRequest(t, http.MethodGet, "/api/sessions/class/fast_track_2020_spring_class_A", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder9.Code)
	var sessions []domains.Session
	if err := json.Unmarshal(recorder9.Body.Bytes(), &sessions); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, sessions[0].Id)
	assert.EqualValues(t, "fast_track_2020_spring_class_A", sessions[0].ClassId)
	assert.EqualValues(t, 1, sessions[1].Id)
	assert.EqualValues(t, "fast_track_2020_spring_class_A", sessions[1].ClassId)
	assert.EqualValues(t, 2, len(sessions))

	resetSessionTables(t)
}

// Test: Create 1 Session, Update it, GetBySessionId()
func TestUpdateSession(t *testing.T) {
	// Create 1 Session
	start := time.Now().UTC()
	end := start.Add(time.Hour)
	prog1 := createProgram("fast_track", "Fast Track", 1, 12, "descript1", 0)
	loc1 := createLocation("loc_1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	semester1 := createSemester("2020_spring", "Spring 2020", 1)
	class1 := createClassUtil("fast_track", "2020_spring", "class_A", "loc_1", "5 pm - 7 pm", 50, 0)
	session1 := createSession("fast_track_2020_spring_class_A", start, end, false, "special lecture from guest")
	body1 := utils.CreateJsonBody(&prog1)
	body2 := utils.CreateJsonBody(&loc1)
	body3 := utils.CreateJsonBody(&semester1)
	body4 := utils.CreateJsonBody(&class1)
	body5 := utils.CreateJsonBody([]domains.Session{session1})
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body3)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body4)
	recorder5 := utils.SendHttpRequest(t, http.MethodPost, "/api/sessions/create", body5)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Update
	updatedSession := createSession("fast_track_2020_spring_class_A", start, end, true, "cancelled due to corona")
	updatedBody := utils.CreateJsonBody(&updatedSession)
	recorder6 := utils.SendHttpRequest(t, http.MethodPost, "/api/sessions/session/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder6.Code)

	// Get
	recorder7 := utils.SendHttpRequest(t, http.MethodGet, "/api/sessions/session/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder7.Code)

	// Validate results
	var session domains.Session
	if err := json.Unmarshal(recorder7.Body.Bytes(), &session); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, session.Id)
	assert.EqualValues(t, "fast_track_2020_spring_class_A", session.ClassId)
	assert.EqualValues(t, domains.NewNullString("cancelled due to corona"), session.Notes)

	resetSessionTables(t)
}

// Test: Create 2 Sessions, Delete them, GetBySessionId()
func TestDeleteSessions(t *testing.T) {
	// Create
	start := time.Now().UTC()
	end := start.Add(time.Hour)
	prog1 := createProgram("fast_track", "Fast Track", 1, 12, "descript1", 0)
	loc1 := createLocation("loc_1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	semester1 := createSemester("2020_spring", "Spring 2020", 1)
	class1 := createClassUtil("fast_track", "2020_spring", "class_A", "loc_1", "5 pm - 7 pm", 50, 0)
	session1 := createSession("fast_track_2020_spring_class_A", start, end, false, "special lecture from guest")
	session2 := createSession("fast_track_2020_spring_class_A", start, end, true, "May 5th regular meeting")
	body1 := utils.CreateJsonBody(&prog1)
	body2 := utils.CreateJsonBody(&loc1)
	body3 := utils.CreateJsonBody(&semester1)
	body4 := utils.CreateJsonBody(&class1)
	body5 := utils.CreateJsonBody([]domains.Session{session1, session2})
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body3)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body4)
	recorder5 := utils.SendHttpRequest(t, http.MethodPost, "/api/sessions/create", body5)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Delete
	body6 := utils.CreateJsonBody([]uint{1, 2})
	recorder6 := utils.SendHttpRequest(t, http.MethodDelete, "/api/sessions/delete", body6)
	assert.EqualValues(t, http.StatusNoContent, recorder6.Code)

	// Get
	recorder7 := utils.SendHttpRequest(t, http.MethodGet, "/api/sessions/session/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder7.Code)
	recorder8 := utils.SendHttpRequest(t, http.MethodGet, "/api/sessions/session/2", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder8.Code)

	resetSessionTables(t)
}

// Helper methods
func createSession(classId string, startsAt time.Time, endsAt time.Time, canceled bool, notes string) domains.Session {
	return domains.Session{
		ClassId:  classId,
		StartsAt: startsAt,
		EndsAt:   endsAt,
		Canceled: canceled,
		Notes:    domains.NewNullString(notes),
	}
}

func createClassUtil(programId, semesterId, classKey, locationId, times string, pricePerSession, priceLump uint) domains.Class {
	return domains.Class{
		ProgramId:       programId,
		SemesterId:      semesterId,
		ClassKey:        domains.NewNullString(classKey),
		LocationId:      locationId,
		Times:           times,
		PricePerSession: domains.NewNullUint(pricePerSession),
		PriceLump:       domains.NewNullUint(priceLump),
	}
}

func resetSessionTables(t *testing.T) {
	utils.ResetTable(t, domains.TABLE_SESSIONS)
	utils.ResetTable(t, domains.TABLE_CLASSES)
	utils.ResetTable(t, domains.TABLE_PROGRAMS)
	utils.ResetTable(t, domains.TABLE_SEMESTERS)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}
