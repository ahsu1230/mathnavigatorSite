package integration_tests

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Test: Create 2 Classes and Publish 1
func Test_PublishClass(t *testing.T) {
	// Create
	createAllProgramsSemestersLocations(t)
	class1 := createClass(1)
	class2 := createClass(2)
	body1 := createJsonBody(&class1)
	body2 := createJsonBody(&class2)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get All Published
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var classes1 []domains.Class
	if err := json.Unmarshal(recorder3.Body.Bytes(), &classes1); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 0, len(classes1))

	// Publish
	classIds := []string{"program1_2020_spring_class2"}
	body3 := createJsonBody(&classIds)
	recorder4 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/publish", body3)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Get All Published
	recorder5 := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Validate results
	var classes2 []domains.Class
	if err := json.Unmarshal(recorder5.Body.Bytes(), &classes2); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertClass(t, 2, classes2[0])
	assert.EqualValues(t, 1, len(classes2))

	// Get All Unpublished
	recorder6 := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)
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

// Test: Create 2 Achievements and Publish 1
func Test_PublishAchievement(t *testing.T) {
	// Create
	achieve1 := createAchievement(2020, "message1")
	achieve2 := createAchievement(2021, "message2")
	body1 := createJsonBody(&achieve1)
	body2 := createJsonBody(&achieve2)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get All Published
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/achievements/v1/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var achieves1 []domains.Achieve
	if err := json.Unmarshal(recorder3.Body.Bytes(), &achieves1); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 0, len(achieves1))

	// Publish
	ids := []uint{1}
	body3 := createJsonBody(&ids)
	recorder4 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/publish", body3)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Get All Published
	recorder5 := sendHttpRequest(t, http.MethodGet, "/api/achievements/v1/all?published=true", nil)
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
	recorder6 := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)
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

	resetTable(t, domains.TABLE_ACHIEVEMENTS)
}
