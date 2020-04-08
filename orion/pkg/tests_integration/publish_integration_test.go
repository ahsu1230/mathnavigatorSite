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
	resetTable(t, domains.TABLE_ACHIEVEMENTS)

	achieve1 := createAchievement(2020, "message1")
	achieve2 := createAchievement(2021, "message2")
	body1 := createJsonBody(&achieve1)
	body2 := createJsonBody(&achieve2)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Call Get Unpublished!
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder3.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, unpublishedDomains.Achieves[0].Id)
	assert.EqualValues(t, 2020, unpublishedDomains.Achieves[0].Year)
	assert.EqualValues(t, "message1", unpublishedDomains.Achieves[0].Message)
	assert.EqualValues(t, 2, unpublishedDomains.Achieves[1].Id)
	assert.EqualValues(t, 2021, unpublishedDomains.Achieves[1].Year)
	assert.EqualValues(t, "message2", unpublishedDomains.Achieves[1].Message)
	assert.EqualValues(t, 2, len(unpublishedDomains.Achieves))
}
