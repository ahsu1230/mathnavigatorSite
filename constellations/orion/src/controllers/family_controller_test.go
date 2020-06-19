package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

//
// Test Get Family
//
func TestGetFamily_Success(t *testing.T) {
	testUtils.FamilyRepo.MockSelectById = func(id uint) (domains.Family, error) {
		family := createMockFamily(
			1,
			"john_smith@example.com",
			"password",
		)
		return family, nil
	}
	repos.FamilyRepo = &familyRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/families/family/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var family domains.Family
	if err := json.Unmarshal(recorder.Body.Bytes(), &family); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, family.Id)
	assert.EqualValues(t, "john_smith@example.com", family.PrimaryEmail)
	assert.EqualValues(t, "password", family.Password)
}

func TestSearchFamily_Success(t *testing.T) {
	testUtils.FamilyRepo.MockSelectByPrimaryEmail = func(primaryEmail string) (domains.Family, error) {
		return createMockFamily(
			1,
			"john_smith@example.com",
			"password",
		), nil
	}
	repos.FamilyRepo = &familyRepo

	// Create search body for HTTP request
	familySearchBody := controllers.FamilySearchBody{
		"john_smith@example.com",
	}
	marshal, err := json.Marshal(&familySearchBody)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBuffer(marshal)

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodPost, "/api/families/search", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var families domains.Family
	if err := json.Unmarshal(recorder.Body.Bytes(), &families); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, families.Id)
	assert.EqualValues(t, "john_smith@example.com", families.PrimaryEmail)
	assert.EqualValues(t, "password", families.Password)
}

func TestGetFamily_Failure(t *testing.T) {
	testUtils.FamilyRepo.MockSelectById = func(id uint) (domains.Family, error) {
		return domains.Family{}, errors.New("not found")
	}
	repos.FamilyRepo = &familyRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/families/family/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateFamily_Success(t *testing.T) {
	testUtils.FamilyRepo.MockInsert = func(family domains.Family) error {
		return nil
	}
	repos.FamilyRepo = &familyRepo

	// Create new HTTP request to endpoint
	family := createMockFamily(
		1,
		"john_smith@example.com",
		"password",
	)
	body := createBodyFromFamily(family)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/families/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateFamily_Failure(t *testing.T) {
	// no mock needed
	repos.FamilyRepo = &familyRepo

	// Create new HTTP request to endpoint
	family := createMockFamily(
		1,
		"",
		"",
	)
	body := createBodyFromFamily(family)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/families/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateFamily_Success(t *testing.T) {
	testUtils.FamilyRepo.MockUpdate = func(id uint, family domains.Family) error {
		return nil // Successful update
	}
	repos.FamilyRepo = &familyRepo

	// Create new HTTP request to endpoint
	family := createMockFamily(
		1,
		"john_smith@example.com",
		"password",
	)
	body := createBodyFromFamily(family)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/families/family/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateFamily_Invalid(t *testing.T) {
	// no mock needed
	repos.FamilyRepo = &familyRepo

	// Create new HTTP request to endpoint
	family := createMockFamily(
		1,
		"",
		"",
	)
	body := createBodyFromFamily(family)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/families/family/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateFamily_Failure(t *testing.T) {
	testUtils.FamilyRepo.MockUpdate = func(id uint, family domains.Family) error {
		return errors.New("not found")
	}
	repos.FamilyRepo = &familyRepo

	// Create new HTTP request to endpoint
	family := createMockFamily(
		1,
		"john_smith@example.com",
		"password",
	)
	body := createBodyFromFamily(family)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/families/family/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteFamily_Success(t *testing.T) {
	testUtils.FamilyRepo.MockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	repos.FamilyRepo = &familyRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/families/family/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteFamily_Failure(t *testing.T) {
	testUtils.FamilyRepo.MockDelete = func(id uint) error {
		return errors.New("not found")
	}
	repos.FamilyRepo = &familyRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/families/family/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockFamily(id uint, primary_email string, password string) domains.Family {
	return domains.Family{
		Id:           id,
		PrimaryEmail: primary_email,
		Password:     password,
	}
}

func createBodyFromFamily(family domains.Family) io.Reader {
	marshal, err := json.Marshal(&family)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
