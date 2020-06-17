package integration_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/stretchr/testify/assert"
)

//create 1 family then get by id 
func Test_SearchFamilyById(t *testing.T) {
	family1 := createFamily(1)
	body1 := createJsonBody(&family1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	recorder2 := sendHttpRequest(t, http.MethodGet, "/api/families/family/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	var family domains.Family
	if err := json.Unmarshal(recorder2.Body.Bytes(), &family); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertFamily(t, 1, family)

	resetTable(t, domains.TABLE_FAMILIES)
}


// Test: Create 3 Families and search by pagination
func Test_SearchFamilyByPrimaryEmail(t *testing.T) {
	family1 := createFamily(1)
 	body1 := createJsonBody(&family1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Create search body for HTTP request
	familySearchBody := controllers.FamilySearchBody{
		"john_smith@example.com",
	}
	marshal, err := json.Marshal(&familySearchBody)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBuffer(marshal)
	
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/families/search",body)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	var family domains.Family
	if err := json.Unmarshal(recorder2.Body.Bytes(), &family); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertFamily(t, 1, family)

	resetTable(t, domains.TABLE_FAMILIES)
}

// Test: Create 3 Families and GetFamilyById
func Test_GetFamilyById(t *testing.T) {
	family1 := createFamily(1)
 	body1 := createJsonBody(&family1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Call Get All!
	recorder := sendHttpRequest(t, http.MethodGet, "/api/families/family/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var family domains.Family
	if err := json.Unmarshal(recorder.Body.Bytes(), &family); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "john_smith@example.com", family.PrimaryEmail)
	assert.EqualValues(t, "password1", family.Password)

	resetTable(t, domains.TABLE_FAMILIES)
}

// Test: Create 1 Family, Update it, GetFamilyById()
func Test_UpdateFamily(t *testing.T) {
	// Create 1 Family
	family1 := createFamily(1)
	body1 := createJsonBody(&family1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedFamily := createFamily(2)
	updatedBody := createJsonBody(&updatedFamily)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/families/family/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/families/family/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var family domains.Family
	if err := json.Unmarshal(recorder3.Body.Bytes(), &family); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertFamily(t, 2, family)

	resetTable(t, domains.TABLE_FAMILIES)
}

// Test: Create 1 Family, Delete it, GetByFamilyId()
func Test_DeleteFamily(t *testing.T) {
	// Create
	family1 := createFamily(1)
	body1 := createJsonBody(&family1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/families/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/families/family/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/families/family/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	resetTable(t, domains.TABLE_FAMILIES)
}

// Helper methods
func createFamily(id int) domains.Family {
	switch id {
	case 1:
		return domains.Family{
			PrimaryEmail:      "john_smith@example.com",
			Password:   "password1",
		}
	case 2:
		return domains.Family{
			PrimaryEmail:      "bob_smith@example.com",
			Password:   "password2",
		}
	case 3:
		return domains.Family{
			PrimaryEmail:      "foobar@example.com",
			Password:   "password3",
		}
	default:
		return domains.Family{}
	}
}

func createAllFamilies(t *testing.T) {
	for i := 1; i < 4; i++ {
		family := createFamily(i)
		body := createJsonBody(&family)
		recorder := sendHttpRequest(t, http.MethodPost, "/api/families/create", body)
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
