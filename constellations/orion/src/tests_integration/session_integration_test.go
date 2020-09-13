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
	utils.SendCreateProgram(t, true, "fast_track", "Fast Track", 1, 12, "descript1", domains.FEATURED_NONE)
	utils.SendCreateProgram(t, true, "slow_track", "Slow Track", 1, 12, "descript1", domains.FEATURED_POPULAR)
	utils.SendCreateLocation(t, true, "loc_1", "Potomac High School", "4040 Location Rd", "City", "MA", "77294", "Room 1", false)
	utils.SendCreateSemester(t, true, domains.SPRING, 2020)
	utils.SendCreateSemester(t, true, domains.FALL, 2020)
	utils.SendCreateClass(t, true, "fast_track", "2020_spring", "class_A", "loc_1", "5 pm - 7 pm", 50, 0)
	utils.SendCreateClass(t, true, "slow_track", "2020_fall", "class_B", "loc_1", "3 pm - 7 pm", 50, 0)
	utils.SendCreateSession(t, true, "fast_track_2020_spring_class_A", mid, end, false, "special lecture from guest")
	utils.SendCreateSession(t, true, "fast_track_2020_spring_class_A", start, end, true, "May 5th regular meeting")
	utils.SendCreateSession(t, true, "slow_track_2020_fall_class_B", start, end, false, "May 5th regular meeting")

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
	utils.SendCreateProgram(t, true, "fast_track", "Fast Track", 1, 12, "descript1", domains.FEATURED_NONE)
	utils.SendCreateLocation(t, true, "loc_1", "Potomac High School", "4040 Location Rd", "City", "MA", "77294", "Room 1", false)
	utils.SendCreateSemester(t, true, domains.SPRING, 2020)
	utils.SendCreateClass(t, true, "fast_track", "2020_spring", "class_A", "loc_1", "5 pm - 7 pm", 50, 0)
	utils.SendCreateSession(t, true, "fast_track_2020_spring_class_A", start, end, false, "special lecture from guest")

	// Update
	updatedSession := domains.Session{
		ClassId:  "fast_track_2020_spring_class_A",
		StartsAt: start,
		EndsAt:   end,
		Canceled: true,
		Notes:    domains.NewNullString("cancelled due to corona"),
	}
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
	utils.SendCreateProgram(t, true, "fast_track", "Fast Track", 1, 12, "descript1", domains.FEATURED_NONE)
	utils.SendCreateLocation(t, true, "loc_1", "Potomac High School", "4040 Location Rd", "City", "MA", "77294", "Room 1", false)
	utils.SendCreateSemester(t, true, domains.SPRING, 2020)
	utils.SendCreateClass(t, true, "fast_track", "2020_spring", "class_A", "loc_1", "5 pm - 7 pm", 50, 0)

	session1 := domains.Session{
		ClassId:  "fast_track_2020_spring_class_A",
		StartsAt: start,
		EndsAt:   end,
		Canceled: false,
		Notes:    domains.NewNullString("special lecture from guest"),
	}
	session2 := domains.Session{
		ClassId:  "fast_track_2020_spring_class_A",
		StartsAt: start,
		EndsAt:   end,
		Canceled: true,
		Notes:    domains.NewNullString("Canceled due to holiday"),
	}
	sessions := []domains.Session{session1, session2}
	utils.SendCreateSessions(t, true, sessions)

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

func resetSessionTables(t *testing.T) {
	utils.ResetTable(t, domains.TABLE_SESSIONS)
	utils.ResetTable(t, domains.TABLE_CLASSES)
	utils.ResetTable(t, domains.TABLE_PROGRAMS)
	utils.ResetTable(t, domains.TABLE_SEMESTERS)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}
