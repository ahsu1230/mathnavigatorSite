package integration_tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Families and GetAll()
func Test_CreateFamilies(t *testing.T) {
	createAllFamilies(t)

	// Call Get All!
	recorder := sendHttpRequest(t, http.MethodPost, "/api/families/create", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var families []domains.Family
	if err := json.Unmarshal(recorder.Body.Bytes(), &families); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertFamily(t, 1, families[0])
	assertFamily(t, 2, families[1])
	assertFamily(t, 3, families[2])
	assert.EqualValues(t, 3, len(families))

	resetTable(t, domains.TABLE_FAMILIES)
}

// Test: Create 3 Families and search by pagination
func Test_SearchFamily(t *testing.T) {
	createAllFamilies(t)

	// Call Get All Searching for "Smith" With Page Size 2 Offset 0
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/families/all?search=Smith&pageSize=2&offset=0", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	var families []domains.Family
	if err := json.Unmarshal(recorder1.Body.Bytes(), &families); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertFamily(t, 1, families[0])
	assertFamily(t, 2, families[1])
	assert.EqualValues(t, 2, len(families))

	// Call Get All Searching for "Smith" With Page Size 2 Offset 2
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/families/all?search=Smith&pageSize=2&offset=2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	var families2 []domains.Family
	if err := json.Unmarshal(recorder2.Body.Bytes(), &families2); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertFamily(t, 3, families2[0])
	assert.EqualValues(t, 1, len(families2))

	resetTable(t, domains.TABLE_FAMILIES)
}

// Test: Create 3 Families and GetFamilyById
func Test_GetFamilyById(t *testing.T) {
	createAllFamilies(t)

	// Call Get All!
	recorder := sendHttpRequest(t, http.MethodGet, "/api/families/family/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var families []domains.Family
	if err := json.Unmarshal(recorder.Body.Bytes(), &families); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertFamily(t, 2, families[0])
	assertFamily(t, 3, families[1])
	assert.EqualValues(t, 2, len(families))

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
