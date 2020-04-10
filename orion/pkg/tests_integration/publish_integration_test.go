package integration_tests

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Test: GetUnpublished()
func Test_GetUnpublished(t *testing.T) {
	resetAllTables(t)

	createAllProgramsSemestersLocations(t)
	class1 := createClass(1)
	class2 := createClass(2)
	classBody1 := createJsonBody(&class1)
	classBody2 := createJsonBody(&class2)
	classRecorder1 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", classBody1)
	classRecorder2 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", classBody2)
	assert.EqualValues(t, http.StatusOK, classRecorder1.Code)
	assert.EqualValues(t, http.StatusOK, classRecorder2.Code)

	achieve1 := createAchievement(2020, "message1")
	achieve2 := createAchievement(2021, "message2")
	achieveBody1 := createJsonBody(&achieve1)
	achieveBody2 := createJsonBody(&achieve2)
	achieveRecorder1 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", achieveBody1)
	achieveRecorder2 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", achieveBody2)
	assert.EqualValues(t, http.StatusOK, achieveRecorder1.Code)
	assert.EqualValues(t, http.StatusOK, achieveRecorder2.Code)

	// Call Get Unpublished!
	recorder := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Validate results
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertClass(t, 1, unpublishedDomains.Classes[0])
	assertClass(t, 2, unpublishedDomains.Classes[1])
	assert.EqualValues(t, 2, len(unpublishedDomains.Classes))

	assert.EqualValues(t, 1, unpublishedDomains.Achieves[0].Id)
	assert.EqualValues(t, 2020, unpublishedDomains.Achieves[0].Year)
	assert.EqualValues(t, "message1", unpublishedDomains.Achieves[0].Message)
	assert.EqualValues(t, 2, unpublishedDomains.Achieves[1].Id)
	assert.EqualValues(t, 2021, unpublishedDomains.Achieves[1].Year)
	assert.EqualValues(t, "message2", unpublishedDomains.Achieves[1].Message)
	assert.EqualValues(t, 2, len(unpublishedDomains.Achieves))
}
