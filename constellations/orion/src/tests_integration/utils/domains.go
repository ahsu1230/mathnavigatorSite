package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// This file contains functions to help with integration tests
// It mainly contains functions for creating domains and checking if the craete request was successful

type CreateResponseBody struct {
	Id uint `json:"id"`
}

// Achievements
func SendCreateAchievement(
	t *testing.T,
	checkOk bool,
	year uint,
	message string,
	position uint,
) (uint, *httptest.ResponseRecorder) {
	achieve := domains.Achieve{
		Year:     year,
		Message:  message,
		Position: position,
	}
	body := CreateJsonBody(&achieve)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

// Announcements
func SendCreateAnnouncement(
	t *testing.T,
	checkOk bool,
	postedAt time.Time,
	author string,
	message string,
	onHomePage bool,
) (uint, *httptest.ResponseRecorder) {
	domain := domains.Announce{
		PostedAt:   postedAt,
		Author:     author,
		Message:    message,
		OnHomePage: onHomePage,
	}
	body := CreateJsonBody(&domain)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/announcements/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

// Ask For Help
func SendCreateAskForHelp(
	t *testing.T,
	checkOk bool,
	startsAt time.Time,
	endsAt time.Time,
	title string,
	subject string,
	locationId string,
	notes string,
) (uint, *httptest.ResponseRecorder) {
	domain := domains.AskForHelp{
		StartsAt:   startsAt,
		EndsAt:     endsAt,
		Title:      title,
		Subject:    subject,
		LocationId: locationId,
		Notes:      domains.NewNullString(notes),
	}
	body := CreateJsonBody(&domain)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

// Classes
func SendCreateClass(
	t *testing.T,
	checkOk bool,
	programId string,
	semesterId string,
	classKey string,
	locationId string,
	times string,
	pricePerSession uint,
	priceLump uint,
) (uint, *httptest.ResponseRecorder) {
	domain := domains.Class{
		ProgramId:       programId,
		SemesterId:      semesterId,
		ClassKey:        domains.NewNullString(classKey),
		LocationId:      locationId,
		TimesStr:        times,
		PricePerSession: domains.NewNullUint(pricePerSession),
		PriceLumpSum:    domains.NewNullUint(priceLump),
	}
	body := CreateJsonBody(&domain)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/classes/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

// Locations
func SendCreateLocation(
	t *testing.T,
	checkOk bool,
	locationId string,
	title string,
	street string,
	city string,
	state string,
	zipcode string,
	room string,
	isOnline bool,
) (uint, *httptest.ResponseRecorder) {
	domain := domains.Location{
		LocationId: locationId,
		Title:      title,
		Street:     domains.NewNullString(street),
		City:       domains.NewNullString(city),
		State:      domains.NewNullString(state),
		Zipcode:    domains.NewNullString(zipcode),
		Room:       domains.NewNullString(room),
		IsOnline:   isOnline,
	}
	body := CreateJsonBody(&domain)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/locations/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

// Programs
func SendCreateProgram(
	t *testing.T,
	checkOk bool,
	programId string,
	name string,
	grade1 uint,
	grade2 uint,
	subject string,
	description string,
	featured string,
) (uint, *httptest.ResponseRecorder) {
	domain := domains.Program{
		ProgramId:   programId,
		Title:       name,
		Grade1:      grade1,
		Grade2:      grade2,
		Subject:     subject,
		Description: description,
		Featured:    featured,
	}
	body := CreateJsonBody(&domain)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/programs/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

// Semesters
func SendCreateSemester(
	t *testing.T,
	checkOk bool,
	season string,
	year uint,
) (uint, *httptest.ResponseRecorder) {
	semesterId := fmt.Sprintf("%d_%s", year, season)
	title := strings.Title(fmt.Sprintf("%s %d", season, year))
	domain := domains.Semester{
		SemesterId: semesterId,
		Season:     season,
		Year:       year,
		Title:      title,
	}
	body := CreateJsonBody(&domain)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

// Sessions
func SendCreateSessions(
	t *testing.T,
	checkOk bool,
	sessions []domains.Session,
) (uint, *httptest.ResponseRecorder) {
	body := CreateJsonBody(&sessions)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/sessions/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

func SendCreateSession(
	t *testing.T,
	checkOk bool,
	classId string,
	startsAt time.Time,
	endsAt time.Time,
	canceled bool,
	notes string,
) (uint, *httptest.ResponseRecorder) {
	session := domains.Session{
		ClassId:  classId,
		StartsAt: startsAt,
		EndsAt:   endsAt,
		Canceled: canceled,
		Notes:    domains.NewNullString(notes),
	}
	return SendCreateSessions(t, checkOk, []domains.Session{session})
}

// Transactions
func SendCreateTransaction(
	t *testing.T,
	checkOk bool,
	accountId uint,
	paymentType string,
	amount int,
	notes string,
) (uint, *httptest.ResponseRecorder) {
	transaction := domains.Transaction{
		AccountId: accountId,
		Type:      paymentType,
		Amount:    amount,
		Notes:     domains.NewNullString(notes),
	}
	body := CreateJsonBody(&transaction)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/transactions/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

// Accounts
func SendCreateAccountUser(
	t *testing.T,
	checkOk bool,
	account domains.Account,
	user domains.User,
) (uint, *httptest.ResponseRecorder) {
	accountUser := domains.AccountUser{
		Account: account,
		User:    user,
	}
	body := CreateJsonBody(&accountUser)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

// Users
func SendCreateUser(
	t *testing.T,
	checkOk bool,
	user domains.User,
) (uint, *httptest.ResponseRecorder) {
	body := CreateJsonBody(&user)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/users/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

// User-AFHs
func SendCreateUserAfh(
	t *testing.T,
	checkOk bool,
	afhId uint,
	userId uint,
	accountId uint,
) (uint, *httptest.ResponseRecorder) {
	domain := domains.UserAfh{
		AfhId:     afhId,
		UserId:    userId,
		AccountId: accountId,
	}
	body := CreateJsonBody(&domain)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/user-afhs/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}

// User-Classes
func SendCreateUserClass(
	t *testing.T,
	checkOk bool,
	classId string,
	userId uint,
	accountId uint,
) (uint, *httptest.ResponseRecorder) {
	domain := domains.UserClass{
		ClassId:   classId,
		UserId:    userId,
		AccountId: accountId,
	}
	body := CreateJsonBody(&domain)
	recorder := SendHttpRequest(t, http.MethodPost, "/api/user-classes/create", body)
	if checkOk {
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
	var responseBody CreateResponseBody
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Unexpected error when creating domain: %v\n", err)
	}
	return responseBody.Id, recorder
}
