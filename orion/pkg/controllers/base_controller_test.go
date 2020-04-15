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

var programService mockProgramService

// Fake programService that implements ProgramService interface
type mockProgramService struct {
	mockGetAll            func(bool) ([]domains.Program, error)
	mockGetAllUnpublished func() ([]domains.Program, error)
	mockGetByProgramId    func(string) (domains.Program, error)
	mockCreate            func(domains.Program) error
	mockUpdate            func(string, domains.Program) error
	mockPublish           func([]string) []domains.PublishErrorBody
	mockDelete            func(string) error
}

// Implement methods of ProgramService interface with mocked implementations
func (programService *mockProgramService) GetAll(publishedOnly bool) ([]domains.Program, error) {
	return programService.mockGetAll(publishedOnly)
}
func (programService *mockProgramService) GetAllUnpublished() ([]domains.Program, error) {
	return programService.mockGetAllUnpublished()
}
func (programService *mockProgramService) GetByProgramId(programId string) (domains.Program, error) {
	return programService.mockGetByProgramId(programId)
}
func (programService *mockProgramService) Create(program domains.Program) error {
	return programService.mockCreate(program)
}
func (programService *mockProgramService) Update(programId string, program domains.Program) error {
	return programService.mockUpdate(programId, program)
}
func (programService *mockProgramService) Publish(programIds []string) []domains.PublishErrorBody {
	return programService.mockPublish(programIds)
}
func (programService *mockProgramService) Delete(programId string) error {
	return programService.mockDelete(programId)
}

var classService mockClassService

// Fake classService that implements ClassService interface
type mockClassService struct {
	mockGetAll                    func(bool) ([]domains.Class, error)
	mockGetAllUnpublished            func() ([]domains.Class, error)
	mockGetByClassId              func(string) (domains.Class, error)
	mockGetByProgramId            func(string) ([]domains.Class, error)
	mockGetBySemesterId           func(string) ([]domains.Class, error)
	mockGetByProgramAndSemesterId func(string, string) ([]domains.Class, error)
	mockCreate                    func(domains.Class) error
	mockUpdate                    func(string, domains.Class) error
	mockPublish                   func([]string) error
	mockDelete                    func(string) error
}

// Implement methods of ClassService interface with mocked implementations
func (classService *mockClassService) GetAll(publishedOnly bool) ([]domains.Class, error) {
	return classService.mockGetAll(publishedOnly)
}
func (classService *mockClassService) GetAllUnpublished() ([]domains.Class, error) {
	return classService.mockGetAllUnpublished()
}
func (classService *mockClassService) GetByClassId(classId string) (domains.Class, error) {
	return classService.mockGetByClassId(classId)
}
func (classService *mockClassService) GetByProgramId(programId string) ([]domains.Class, error) {
	return classService.mockGetByProgramId(programId)
}
func (classService *mockClassService) GetBySemesterId(semesterId string) ([]domains.Class, error) {
	return classService.mockGetBySemesterId(semesterId)
}
func (classService *mockClassService) GetByProgramAndSemesterId(programId, semesterId string) ([]domains.Class, error) {
	return classService.mockGetByProgramAndSemesterId(programId, semesterId)
}
func (classService *mockClassService) Create(class domains.Class) error {
	return classService.mockCreate(class)
}
func (classService *mockClassService) Update(classId string, class domains.Class) error {
	return classService.mockUpdate(classId, class)
}
func (classService *mockClassService) Publish(classIds []string) error {
	return classService.mockPublish(classIds)
}
func (classService *mockClassService) Delete(classId string) error {
	return classService.mockDelete(classId)
}

var locationService mockLocationService

// Fake locationService that implements LocationService interface
type mockLocationService struct {
	mockGetAll            func(bool) ([]domains.Location, error)
	mockGetAllUnpublished func() ([]domains.Location, error)
	mockGetByLocationId   func(string) (domains.Location, error)
	mockCreate            func(domains.Location) error
	mockUpdate            func(string, domains.Location) error
	mockPublish           func([]string) []domains.PublishErrorBody
	mockDelete            func(string) error
}

// Implement methods of LocationService interface with mocked implementations
func (locationService *mockLocationService) GetAll(publishedOnly bool) ([]domains.Location, error) {
	return locationService.mockGetAll(publishedOnly)
}
func (locationService *mockLocationService) GetAllUnpublished() ([]domains.Location, error) {
	return locationService.mockGetAllUnpublished()
}
func (locationService *mockLocationService) GetByLocationId(locId string) (domains.Location, error) {
	return locationService.mockGetByLocationId(locId)
}
func (locationService *mockLocationService) Create(location domains.Location) error {
	return locationService.mockCreate(location)
}
func (locationService *mockLocationService) Update(locId string, location domains.Location) error {
	return locationService.mockUpdate(locId, location)
}
func (locationService *mockLocationService) Publish(locIds []string) []domains.PublishErrorBody {
	return locationService.mockPublish(locIds)
}
func (locationService *mockLocationService) Delete(locId string) error {
	return locationService.mockDelete(locId)
}

var announceService mockAnnounceService

// Fake announceService that implements AnnounceService interface
type mockAnnounceService struct {
	mockGetAll          func() ([]domains.Announce, error)
	mockGetByAnnounceId func(uint) (domains.Announce, error)
	mockCreate          func(domains.Announce) error
	mockUpdate          func(uint, domains.Announce) error
	mockDelete          func(uint) error
}

// Implement methods of AnnounceService interface with mocked implementations
func (announceService *mockAnnounceService) GetAll() ([]domains.Announce, error) {
	return announceService.mockGetAll()
}
func (announceService *mockAnnounceService) GetByAnnounceId(id uint) (domains.Announce, error) {
	return announceService.mockGetByAnnounceId(id)
}
func (announceService *mockAnnounceService) Create(announce domains.Announce) error {
	return announceService.mockCreate(announce)
}
func (announceService *mockAnnounceService) Update(id uint, announce domains.Announce) error {
	return announceService.mockUpdate(id, announce)
}
func (announceService *mockAnnounceService) Delete(id uint) error {
	return announceService.mockDelete(id)
}

var achieveService mockAchieveService

// Fake achieveService that implements AchieveService interface
type mockAchieveService struct {
	mockGetAll              func(bool) ([]domains.Achieve, error)
	mockGetAllUnpublished      func() ([]domains.Achieve, error)
	mockGetById             func(uint) (domains.Achieve, error)
	mockGetAllGroupedByYear func() ([]domains.AchieveYearGroup, error)
	mockCreate              func(domains.Achieve) error
	mockUpdate              func(uint, domains.Achieve) error
	mockPublish             func([]uint) error
	mockDelete              func(uint) error
}

// Implement methods of AchieveService interface with mocked implementations
func (achieveService *mockAchieveService) GetAll(publishedOnly bool) ([]domains.Achieve, error) {
	return achieveService.mockGetAll(publishedOnly)
}
func (achieveService *mockAchieveService) GetAllUnpublished() ([]domains.Achieve, error) {
	return achieveService.mockGetAllUnpublished()
}
func (achieveService *mockAchieveService) GetById(id uint) (domains.Achieve, error) {
	return achieveService.mockGetById(id)
}
func (achieveService *mockAchieveService) GetAllGroupedByYear() ([]domains.AchieveYearGroup, error) {
	return achieveService.mockGetAllGroupedByYear()
}
func (achieveService *mockAchieveService) Create(achieve domains.Achieve) error {
	return achieveService.mockCreate(achieve)
}
func (achieveService *mockAchieveService) Update(id uint, achieve domains.Achieve) error {
	return achieveService.mockUpdate(id, achieve)
}
func (achieveService *mockAchieveService) Publish(ids []uint) error {
	return achieveService.mockPublish(ids)
}
func (achieveService *mockAchieveService) Delete(id uint) error {
	return achieveService.mockDelete(id)
}

var semesterService mockSemesterService

// Fake semesterService that implements SemesterService interface
type mockSemesterService struct {
	mockGetAll          func(bool) ([]domains.Semester, error)
	mockGetAllUnpublished  func() ([]domains.Semester, error)
	mockGetBySemesterId func(string) (domains.Semester, error)
	mockCreate          func(domains.Semester) error
	mockUpdate          func(string, domains.Semester) error
	mockPublish         func([]string) error
	mockDelete          func(string) error
}

// Implement methods of SemesterService interface with mocked implementations
func (semesterService *mockSemesterService) GetAll(publishedOnly bool) ([]domains.Semester, error) {
	return semesterService.mockGetAll(publishedOnly)
}
func (semesterService *mockSemesterService) GetAllUnpublished() ([]domains.Semester, error) {
	return semesterService.mockGetAllUnpublished()
}
func (semesterService *mockSemesterService) GetBySemesterId(semesterId string) (domains.Semester, error) {
	return semesterService.mockGetBySemesterId(semesterId)
}
func (semesterService *mockSemesterService) Create(semester domains.Semester) error {
	return semesterService.mockCreate(semester)
}
func (semesterService *mockSemesterService) Update(semesterId string, semester domains.Semester) error {
	return semesterService.mockUpdate(semesterId, semester)
}
func (semesterService *mockSemesterService) Publish(semesterIds []string) error {
	return semesterService.mockPublish(semesterIds)
}
func (semesterService *mockSemesterService) Delete(semesterId string) error {
	return semesterService.mockDelete(semesterId)
}

var sessionService mockSessionService

// Fake sessionService that implements SessionService interface
type mockSessionService struct {
	mockGetAllByClassId   func(string, bool) ([]domains.Session, error)
	mockGetAllUnpublished func() ([]domains.Session, error)
	mockGetBySessionId    func(uint) (domains.Session, error)
	mockCreate            func(domains.Session) error
	mockUpdate            func(uint, domains.Session) error
	mockPublish           func([]uint) []domains.PublishErrorBody
	mockDelete            func(uint) error
}

// Implement methods of SessionService interface with mocked implementations
func (sessionService *mockSessionService) GetAllByClassId(classId string, publishedOnly bool) ([]domains.Session, error) {
	return sessionService.mockGetAllByClassId(classId, publishedOnly)
}
func (sessionService *mockSessionService) GetAllUnpublished() ([]domains.Session, error) {
	return sessionService.mockGetAllUnpublished()
}
func (sessionService *mockSessionService) GetBySessionId(id uint) (domains.Session, error) {
	return sessionService.mockGetBySessionId(id)
}
func (sessionService *mockSessionService) Create(session domains.Session) error {
	return sessionService.mockCreate(session)
}
func (sessionService *mockSessionService) Update(id uint, session domains.Session) error {
	return sessionService.mockUpdate(id, session)
}
func (sessionService *mockSessionService) Publish(ids []uint) []domains.PublishErrorBody {
	return sessionService.mockPublish(ids)
}
func (sessionService *mockSessionService) Delete(id uint) error {
	return sessionService.mockDelete(id)
}

var userService mockUserService

// Fake userService that implements UserService interface
type mockUserService struct {
	mockGetAll          func(string, int, int) ([]domains.User, error)
	mockGetById         func(uint) (domains.User, error)
	mockGetByGuardianId func(uint) ([]domains.User, error)
	mockCreate          func(domains.User) error
	mockUpdate          func(uint, domains.User) error
	mockDelete          func(uint) error
}

// Implement methods of UserService interface with mocked implementations
func (userService *mockUserService) GetAll(search string, pageSize, offset int) ([]domains.User, error) {
	return userService.mockGetAll(search, pageSize, offset)
}
func (userService *mockUserService) GetById(id uint) (domains.User, error) {
	return userService.mockGetById(id)
}
func (userService *mockUserService) GetByGuardianId(guardianId uint) ([]domains.User, error) {
	return userService.mockGetByGuardianId(guardianId)
}
func (userService *mockUserService) Create(user domains.User) error {
	return userService.mockCreate(user)
}
func (userService *mockUserService) Update(id uint, user domains.User) error {
	return userService.mockUpdate(id, user)
}
func (userService *mockUserService) Delete(id uint) error {
	return userService.mockDelete(id)
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
