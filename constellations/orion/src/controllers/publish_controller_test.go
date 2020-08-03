package controllers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

func setupMock() {
	testUtils.ClassRepo.MockSelectAllUnpublished = func() ([]domains.Class, error) {
		return []domains.Class{
			testUtils.CreateMockClass(
				"prog1",
				"2020_fall",
				"classA",
				"churchill",
				"3 pm - 5 pm",
			),
			testUtils.CreateMockClass(
				"prog1",
				"2020_fall",
				"classB",
				"churchill",
				"3 pm - 5 pm",
			),
		}, nil
	}
	testUtils.AchieveRepo.MockSelectAllUnpublished = func() ([]domains.Achieve, error) {
		return []domains.Achieve{
			testUtils.CreateMockAchievement(1, 2020, "message1"),
			testUtils.CreateMockAchievement(2, 2021, "message2"),
		}, nil
	}
	repos.ClassRepo = &testUtils.ClassRepo
	repos.AchieveRepo = &testUtils.AchieveRepo
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

	class0 := unpublishedDomains.Classes[0]
	class1 := unpublishedDomains.Classes[1]
	assert.EqualValues(t, "prog1", class0.ProgramId)
	assert.EqualValues(t, "2020_fall", class0.SemesterId)
	assert.EqualValues(t, "prog1_2020_fall_classA", class0.ClassId)
	assert.EqualValues(t, "prog1", class1.ProgramId)
	assert.EqualValues(t, "2020_fall", class1.SemesterId)
	assert.EqualValues(t, "prog1_2020_fall_classB", class1.ClassId)
	assert.EqualValues(t, 2, len(unpublishedDomains.Classes))

	assert.EqualValues(t, 1, unpublishedDomains.Achieves[0].Id)
	assert.EqualValues(t, 2020, unpublishedDomains.Achieves[0].Year)
	assert.EqualValues(t, "message1", unpublishedDomains.Achieves[0].Message)
	assert.EqualValues(t, 2, unpublishedDomains.Achieves[1].Id)
	assert.EqualValues(t, 2021, unpublishedDomains.Achieves[1].Year)
	assert.EqualValues(t, "message2", unpublishedDomains.Achieves[1].Message)
	assert.EqualValues(t, 2, len(unpublishedDomains.Achieves))
}
