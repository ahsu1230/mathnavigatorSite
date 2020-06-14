package controllers_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

func setupMock() {
	testUtils.ProgramRepo.MockSelectAllUnpublished = func() ([]domains.Program, error) {
		return []domains.Program{
			createMockProgram("prog1", "Program1", 2, 3, "descript1", 0),
			createMockProgram("prog2", "Program2", 8, 12, "descript2", 0),
		}, nil
	}
	testUtils.ClassRepo.MockSelectAllUnpublished = func() ([]domains.Class, error) {
		return createMockClasses(1, 2), nil
	}
	testUtils.LocationRepo.MockSelectAllUnpublished = func() ([]domains.Location, error) {
		return []domains.Location{
			createMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1"),
			createMockLocation("loc2", "4040 Sesame St", "City", "MD", "77294", "Room 2"),
		}, nil
	}
	testUtils.AchieveRepo.MockSelectAllUnpublished = func() ([]domains.Achieve, error) {
		return []domains.Achieve{
			createMockAchievement(1, 2020, "message1"),
			createMockAchievement(2, 2021, "message2"),
		}, nil
	}
	testUtils.SemesterRepo.MockSelectAllUnpublished = func() ([]domains.Semester, error) {
		return []domains.Semester{
			createMockSemester("2020_fall", "Fall 2020"),
			createMockSemester("2020_winter", "Winter 2020"),
		}, nil
	}
	testUtils.SessionRepo.MockSelectAllUnpublished = func() ([]domains.Session, error) {
		now := time.Now().UTC()
		return []domains.Session{
			createMockSession(1, "id_1", now, now, true, "special lecture from guest"),
			createMockSession(2, "id_2", now, now, false, "daily meeting"),
		}, nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo
	repos.ClassRepo = &testUtils.ClassRepo
	repos.LocationRepo = &testUtils.LocationRepo
	repos.AchieveRepo = &testUtils.AchieveRepo
	repos.SemesterRepo = &testUtils.SemesterRepo
	repos.SessionRepo = &testUtils.SessionRepo
}

//
// Test Get Unpublished
//
func TestGetAllUnpublished_Success(t *testing.T) {
	setupMock()

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, "prog1", unpublishedDomains.Programs[0].ProgramId)
	assert.EqualValues(t, "Program1", unpublishedDomains.Programs[0].Name)
	assert.EqualValues(t, "prog2", unpublishedDomains.Programs[1].ProgramId)
	assert.EqualValues(t, "Program2", unpublishedDomains.Programs[1].Name)
	assert.EqualValues(t, 2, len(unpublishedDomains.Programs))

	assertMockClasses(t, 1, unpublishedDomains.Classes[0])
	assertMockClasses(t, 2, unpublishedDomains.Classes[1])
	assert.EqualValues(t, 2, len(unpublishedDomains.Classes))

	assert.EqualValues(t, "loc1", unpublishedDomains.Locations[0].LocationId)
	assert.EqualValues(t, "4040 Location Rd", unpublishedDomains.Locations[0].Street)
	assert.EqualValues(t, "MA", unpublishedDomains.Locations[0].State)
	assert.EqualValues(t, "loc2", unpublishedDomains.Locations[1].LocationId)
	assert.EqualValues(t, "4040 Sesame St", unpublishedDomains.Locations[1].Street)
	assert.EqualValues(t, "MD", unpublishedDomains.Locations[1].State)
	assert.EqualValues(t, 2, len(unpublishedDomains.Locations))

	assert.EqualValues(t, 1, unpublishedDomains.Achieves[0].Id)
	assert.EqualValues(t, 2020, unpublishedDomains.Achieves[0].Year)
	assert.EqualValues(t, "message1", unpublishedDomains.Achieves[0].Message)
	assert.EqualValues(t, 2, unpublishedDomains.Achieves[1].Id)
	assert.EqualValues(t, 2021, unpublishedDomains.Achieves[1].Year)
	assert.EqualValues(t, "message2", unpublishedDomains.Achieves[1].Message)
	assert.EqualValues(t, 2, len(unpublishedDomains.Achieves))

	assert.EqualValues(t, "2020_fall", unpublishedDomains.Semesters[0].SemesterId)
	assert.EqualValues(t, "Fall 2020", unpublishedDomains.Semesters[0].Title)
	assert.EqualValues(t, "2020_winter", unpublishedDomains.Semesters[1].SemesterId)
	assert.EqualValues(t, "Winter 2020", unpublishedDomains.Semesters[1].Title)
	assert.EqualValues(t, 2, len(unpublishedDomains.Semesters))

	assert.EqualValues(t, 1, unpublishedDomains.Sessions[0].Id)
	assert.EqualValues(t, "id_1", unpublishedDomains.Sessions[0].ClassId)
	assert.EqualValues(t, 2, unpublishedDomains.Sessions[1].Id)
	assert.EqualValues(t, "id_2", unpublishedDomains.Sessions[1].ClassId)
	assert.EqualValues(t, 2, len(unpublishedDomains.Sessions))
}
