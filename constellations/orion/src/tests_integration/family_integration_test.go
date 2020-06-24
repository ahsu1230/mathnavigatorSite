package tests_integration

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

//create 1 family then get by id
func Test_SearchFamilyById(t *testing.T) {
	family1 := createFamily(1)
	body1 := utils.CreateJsonBody(&family1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	recorder2 := utils.SendHttpRequest(t, http.MethodGet, "/api/families/family/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	var family domains.Family
	if err := json.Unmarshal(recorder2.Body.Bytes(), &family); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertFamily(t, 1, family)

	utils.ResetTable(t, domains.TABLE_FAMILIES)
}

// Test: Create 3 Families and search by pagination
func Test_SearchFamilyByPrimaryEmail(t *testing.T) {
	family1 := createFamily(1)
	body1 := utils.CreateJsonBody(&family1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	body := strings.NewReader(`{
		"primaryEmail": "john_smith@example.com"
	}`)

	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/families/search", body)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	var family domains.Family
	if err := json.Unmarshal(recorder2.Body.Bytes(), &family); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertFamily(t, 1, family)

	utils.ResetTable(t, domains.TABLE_FAMILIES)
}

// Test: Create 3 Families and GetFamilyById
func Test_GetFamilyById(t *testing.T) {
	family1 := createFamily(1)
	body1 := utils.CreateJsonBody(&family1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Call Get All!
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/families/family/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var family domains.Family
	if err := json.Unmarshal(recorder.Body.Bytes(), &family); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "john_smith@example.com", family.PrimaryEmail)
	assert.EqualValues(t, "password1", family.Password)

	utils.ResetTable(t, domains.TABLE_FAMILIES)
}

// Test: Create 1 Family, Update it, GetFamilyById()
func Test_UpdateFamily(t *testing.T) {
	// Create 1 Family
	family1 := createFamily(1)
	body1 := utils.CreateJsonBody(&family1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedFamily := createFamily(2)
	updatedBody := utils.CreateJsonBody(&updatedFamily)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/families/family/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/families/family/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var family domains.Family
	if err := json.Unmarshal(recorder3.Body.Bytes(), &family); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertFamily(t, 2, family)

	utils.ResetTable(t, domains.TABLE_FAMILIES)
}

// Test: Create 1 Family, Delete it, GetByFamilyId()
func Test_DeleteFamily(t *testing.T) {
	// Create
	family1 := createFamily(1)
	body1 := utils.CreateJsonBody(&family1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/families/family/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/families/family/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_FAMILIES)
}

// Helper methods
func createFamily(id int) domains.Family {
	switch id {
	case 1:
		return domains.Family{
			PrimaryEmail: "john_smith@example.com",
			Password:     "password1",
		}
	case 2:
		return domains.Family{
			PrimaryEmail: "bob_smith@example.com",
			Password:     "password2",
		}
	case 3:
		return domains.Family{
			PrimaryEmail: "foobar@example.com",
			Password:     "password3",
		}
	default:
		return domains.Family{}
	}
}

func createAllFamilies(t *testing.T) {
	for i := 1; i < 4; i++ {
		family := createFamily(i)
		body := utils.CreateJsonBody(&family)
		recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/families/create", body)
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
}

func assertFamily(t *testing.T, id int, family domains.Family) {
	switch id {
	case 1:
		assert.EqualValues(t, "john_smith@example.com", family.PrimaryEmail)
		assert.EqualValues(t, "password1", family.Password)
	case 2:
		assert.EqualValues(t, "bob_smith@example.com", family.PrimaryEmail)
		assert.EqualValues(t, "password2", family.Password)
	case 3:
		assert.EqualValues(t, "foobar@example.com", family.PrimaryEmail)
		assert.EqualValues(t, "password2", family.Password)
	}
}