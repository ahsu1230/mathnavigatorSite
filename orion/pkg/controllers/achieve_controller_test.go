package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

//
// Test Get All
//
func TestGetAllAchievements_Success(t *testing.T) {
	achieveService.mockGetAll = func() ([]domains.Achieve, error) {
		return []domains.Achieve{
			{
				Id:      1,
				Year:    2020,
				Message: "message1",
			},
			{
				Id:      2,
				Year:    2021,
				Message: "message2",
			},
		}, nil
	}
	services.AchieveService = &achieveService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/achievements/v1/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var achieves []domains.Achieve
	if err := json.Unmarshal(recorder.Body.Bytes(), &achieves); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieves[0].Id)
	assert.EqualValues(t, 2020, achieves[0].Year)
	assert.EqualValues(t, "message1", achieves[0].Message)
	assert.EqualValues(t, 2, achieves[1].Id)
	assert.EqualValues(t, 2021, achieves[1].Year)
	assert.EqualValues(t, "message2", achieves[1].Message)
	assert.EqualValues(t, 2, len(achieves))
}

//
// Test Get Achieve
//
func TestGetAchievement_Success(t *testing.T) {
	achieveService.mockGetById = func(id uint) (domains.Achieve, error) {
		achieve := createMockAchievement(1, 2020, "message1")
		return achieve, nil
	}
	services.AchieveService = &achieveService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/achievements/v1/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var achieve domains.Achieve
	if err := json.Unmarshal(recorder.Body.Bytes(), &achieve); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieve.Id)
	assert.EqualValues(t, 2020, achieve.Year)
	assert.EqualValues(t, "message1", achieve.Message)
}

func TestGetAchievement_Failure(t *testing.T) {
	achieveService.mockGetById = func(id uint) (domains.Achieve, error) {
		return domains.Achieve{}, errors.New("not found")
	}
	services.AchieveService = &achieveService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/achievements/v1/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateAchievement_Success(t *testing.T) {
	achieveService.mockCreate = func(achieve domains.Achieve) error {
		return nil
	}
	services.AchieveService = &achieveService

	// Create new HTTP request to endpoint
	achieve := createMockAchievement(1, 2020, "message1")
	marshal, _ := json.Marshal(achieve)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateAchievement_Failure(t *testing.T) {
	// no mock needed
	services.AchieveService = &achieveService

	// Create new HTTP request to endpoint
	achieve := createMockAchievement(1, 0, "")
	marshal, _ := json.Marshal(achieve)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateAchievement_Success(t *testing.T) {
	achieveService.mockUpdate = func(id uint, achieve domains.Achieve) error {
		return nil // Successful update
	}
	services.AchieveService = &achieveService

	// Create new HTTP request to endpoint
	achieve := createMockAchievement(1, 2020, "message1")
	body := createBodyFromAchieve(achieve)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/achievement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateAchievement_Invalid(t *testing.T) {
	// no mock needed
	services.AchieveService = &achieveService

	// Create new HTTP request to endpoint
	achieve := createMockAchievement(1, 0, "")
	body := createBodyFromAchieve(achieve)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/achievement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateAchievement_Failure(t *testing.T) {
	achieveService.mockUpdate = func(id uint, achieve domains.Achieve) error {
		return errors.New("not found")
	}
	services.AchieveService = &achieveService

	// Create new HTTP request to endpoint
	achieve := createMockAchievement(1, 2020, "message1")
	body := createBodyFromAchieve(achieve)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/achievement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteAchievement_Success(t *testing.T) {
	achieveService.mockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	services.AchieveService = &achieveService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/achievements/v1/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteAchievement_Failure(t *testing.T) {
	achieveService.mockDelete = func(id uint) error {
		return errors.New("not found")
	}
	services.AchieveService = &achieveService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/achievements/v1/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockAchievement(id uint, year uint, message string) domains.Achieve {
	return domains.Achieve{
		Id:      id,
		Year:    year,
		Message: message,
	}
}

func createBodyFromAchieve(achieve domains.Achieve) io.Reader {
	marshal, err := json.Marshal(achieve)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
