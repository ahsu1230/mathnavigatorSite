package controllers_test

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/router"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Global test variables
var handler router.Handler

var mps mockProgramService

// Fake programService that implements ProgramService interface
type mockProgramService struct {
	mockGetAll         func() ([]domains.Program, error)
	mockGetByProgramId func(string) (domains.Program, error)
	mockCreate         func(domains.Program) error
	mockUpdate         func(string, domains.Program) error
	mockDelete         func(string) error
}

// Implement methods of ProgramService interface with mocked implementations
func (mps *mockProgramService) GetAll() ([]domains.Program, error) {
	return mps.mockGetAll()
}
func (mps *mockProgramService) GetByProgramId(programId string) (domains.Program, error) {
	return mps.mockGetByProgramId(programId)
}
func (mps *mockProgramService) Create(program domains.Program) error {
	return mps.mockCreate(program)
}
func (mps *mockProgramService) Update(programId string, program domains.Program) error {
	return mps.mockUpdate(programId, program)
}
func (mps *mockProgramService) Delete(programId string) error {
	return mps.mockDelete(programId)
}

var mas mockAnnounceService

// Fake announceService that implements AnnounceService interface
type mockAnnounceService struct {
	mockGetAll          func() ([]domains.Announce, error)
	mockGetByAnnounceId func(uint) (domains.Announce, error)
	mockCreate          func(domains.Announce) error
	mockUpdate          func(uint, domains.Announce) error
	mockDelete          func(uint) error
}

// Implement methods of AnnounceService interface with mocked implementations
func (mas *mockAnnounceService) GetAll() ([]domains.Announce, error) {
	return mas.mockGetAll()
}
func (mas *mockAnnounceService) GetByAnnounceId(id uint) (domains.Announce, error) {
	return mas.mockGetByAnnounceId(id)
}
func (mas *mockAnnounceService) Create(announce domains.Announce) error {
	return mas.mockCreate(announce)
}
func (mas *mockAnnounceService) Update(id uint, announce domains.Announce) error {
	return mas.mockUpdate(id, announce)
}
func (mas *mockAnnounceService) Delete(id uint) error {
	return mas.mockDelete(id)
}

var mls mockLocationService

// Fake locationService that implements LocationService interface
type mockLocationService struct {
	mockGetAll          func() ([]domains.Location, error)
	mockGetByLocationId func(string) (domains.Location, error)
	mockCreate          func(domains.Location) error
	mockUpdate          func(string, domains.Location) error
	mockDelete          func(string) error
}

// Implement methods of LocationService interface with mocked implementations
func (mls *mockLocationService) GetAll() ([]domains.Location, error) {
	return mls.mockGetAll()
}
func (mls *mockLocationService) GetByLocationId(locId string) (domains.Location, error) {
	return mls.mockGetByLocationId(locId)
}
func (mls *mockLocationService) Create(Location domains.Location) error {
	return mls.mockCreate(Location)
}
func (mls *mockLocationService) Update(locId string, Location domains.Location) error {
	return mls.mockUpdate(locId, Location)
}
func (mls *mockLocationService) Delete(locId string) error {
	return mls.mockDelete(locId)
}

func init() {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	handler = router.Handler{Engine: engine}
	handler.SetupApiEndpoints()
}

func sendHttpRequest(t *testing.T, method, url string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Errorf("http request error: %v\n", err)
	}
	w := httptest.NewRecorder()
	handler.Engine.ServeHTTP(w, req)
	return w
}
