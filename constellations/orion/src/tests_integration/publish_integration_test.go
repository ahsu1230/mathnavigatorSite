package tests_integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create 2 Classes and Publish 1
func TestPublishClasses(t *testing.T) {
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
	assert.EqualValues(t, http.StatusNoContent, recorder4.Code)

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
