package controllers_test

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/router"
	"github.com/gin-gonic/gin"
)

// Global test variables
var handler router.Handler

// Utility methods
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

// Mocked structs
var programRepo mockProgramRepo
var classRepo mockClassRepo
var locationRepo mockLocationRepo
var announceRepo mockAnnounceRepo
var achieveRepo mockAchieveRepo
var semesterRepo mockSemesterRepo
var sessionRepo mockSessionRepo
var userRepo mockUserRepo
var familyRepo mockFamilyRepo

// Fake programRepo that implements ProgramRepo interface
type mockProgramRepo struct {
	mockInitialize           func(*sql.DB)
	mockSelectAll            func(bool) ([]domains.Program, error)
	mockSelectAllUnpublished func() ([]domains.Program, error)
	mockSelectByProgramId    func(string) (domains.Program, error)
	mockInsert               func(domains.Program) error
	mockUpdate               func(string, domains.Program) error
	mockPublish              func([]string) error
	mockDelete               func(string) error
}

// Implement methods of ProgramRepo interface with mocked implementations
func (programRepo *mockProgramRepo) Initialize(db *sql.DB) {}
func (programRepo *mockProgramRepo) SelectAll(publishedOnly bool) ([]domains.Program, error) {
	return programRepo.mockSelectAll(publishedOnly)
}
func (programRepo *mockProgramRepo) SelectAllUnpublished() ([]domains.Program, error) {
	return programRepo.mockSelectAllUnpublished()
}
func (programRepo *mockProgramRepo) SelectByProgramId(programId string) (domains.Program, error) {
	return programRepo.mockSelectByProgramId(programId)
}
func (programRepo *mockProgramRepo) Insert(program domains.Program) error {
	return programRepo.mockInsert(program)
}
func (programRepo *mockProgramRepo) Update(programId string, program domains.Program) error {
	return programRepo.mockUpdate(programId, program)
}
func (programRepo *mockProgramRepo) Publish(programIds []string) error {
	return programRepo.mockPublish(programIds)
}
func (programRepo *mockProgramRepo) Delete(programId string) error {
	return programRepo.mockDelete(programId)
}

// Fake classRepo that implements ClassRepo interface
type mockClassRepo struct {
	mockInitialize                   func(*sql.DB)
	mockSelectAll                    func(bool) ([]domains.Class, error)
	mockSelectAllUnpublished         func() ([]domains.Class, error)
	mockSelectByClassId              func(string) (domains.Class, error)
	mockSelectByProgramId            func(string) ([]domains.Class, error)
	mockSelectBySemesterId           func(string) ([]domains.Class, error)
	mockSelectByProgramAndSemesterId func(string, string) ([]domains.Class, error)
	mockInsert                       func(domains.Class) error
	mockUpdate                       func(string, domains.Class) error
	mockPublish                      func([]string) error
	mockDelete                       func(string) error
}

// Implement methods of ClassRepo interface with mocked implementations
func (classRepo *mockClassRepo) Initialize(db *sql.DB) {}
func (classRepo *mockClassRepo) SelectAll(publishedOnly bool) ([]domains.Class, error) {
	return classRepo.mockSelectAll(publishedOnly)
}
func (classRepo *mockClassRepo) SelectAllUnpublished() ([]domains.Class, error) {
	return classRepo.mockSelectAllUnpublished()
}
func (classRepo *mockClassRepo) SelectByClassId(classId string) (domains.Class, error) {
	return classRepo.mockSelectByClassId(classId)
}
func (classRepo *mockClassRepo) SelectByProgramId(programId string) ([]domains.Class, error) {
	return classRepo.mockSelectByProgramId(programId)
}
func (classRepo *mockClassRepo) SelectBySemesterId(semesterId string) ([]domains.Class, error) {
	return classRepo.mockSelectBySemesterId(semesterId)
}
func (classRepo *mockClassRepo) SelectByProgramAndSemesterId(programId, semesterId string) ([]domains.Class, error) {
	return classRepo.mockSelectByProgramAndSemesterId(programId, semesterId)
}
func (classRepo *mockClassRepo) Insert(class domains.Class) error {
	return classRepo.mockInsert(class)
}
func (classRepo *mockClassRepo) Update(classId string, class domains.Class) error {
	return classRepo.mockUpdate(classId, class)
}
func (classRepo *mockClassRepo) Publish(classIds []string) error {
	return classRepo.mockPublish(classIds)
}
func (classRepo *mockClassRepo) Delete(classId string) error {
	return classRepo.mockDelete(classId)
}

// Fake locationRepo that implements LocationRepo interface
type mockLocationRepo struct {
	mockInitialize           func(*sql.DB)
	mockSelectAll            func(bool) ([]domains.Location, error)
	mockSelectAllUnpublished func() ([]domains.Location, error)
	mockSelectByLocationId   func(string) (domains.Location, error)
	mockInsert               func(domains.Location) error
	mockUpdate               func(string, domains.Location) error
	mockPublish              func([]string) error
	mockDelete               func(string) error
}

// Implement methods of LocationRepo interface with mocked implementations
func (locationRepo *mockLocationRepo) Initialize(db *sql.DB) {}
func (locationRepo *mockLocationRepo) SelectAll(publishedOnly bool) ([]domains.Location, error) {
	return locationRepo.mockSelectAll(publishedOnly)
}
func (locationRepo *mockLocationRepo) SelectAllUnpublished() ([]domains.Location, error) {
	return locationRepo.mockSelectAllUnpublished()
}
func (locationRepo *mockLocationRepo) SelectByLocationId(locationId string) (domains.Location, error) {
	return locationRepo.mockSelectByLocationId(locationId)
}
func (locationRepo *mockLocationRepo) Insert(location domains.Location) error {
	return locationRepo.mockInsert(location)
}
func (locationRepo *mockLocationRepo) Update(locationId string, location domains.Location) error {
	return locationRepo.mockUpdate(locationId, location)
}
func (locationRepo *mockLocationRepo) Publish(locationIds []string) error {
	return locationRepo.mockPublish(locationIds)
}
func (locationRepo *mockLocationRepo) Delete(locationId string) error {
	return locationRepo.mockDelete(locationId)
}

// Fake announceRepo that implements AnnounceRepo interface
type mockAnnounceRepo struct {
	mockInitialize         func(*sql.DB)
	mockSelectAll          func() ([]domains.Announce, error)
	mockSelectByAnnounceId func(uint) (domains.Announce, error)
	mockInsert             func(domains.Announce) error
	mockUpdate             func(uint, domains.Announce) error
	mockDelete             func(uint) error
}

// Implement methods of AnnounceRepo interface with mocked implementations
func (announceRepo *mockAnnounceRepo) Initialize(db *sql.DB) {}
func (announceRepo *mockAnnounceRepo) SelectAll() ([]domains.Announce, error) {
	return announceRepo.mockSelectAll()
}
func (announceRepo *mockAnnounceRepo) SelectByAnnounceId(id uint) (domains.Announce, error) {
	return announceRepo.mockSelectByAnnounceId(id)
}
func (announceRepo *mockAnnounceRepo) Insert(announce domains.Announce) error {
	return announceRepo.mockInsert(announce)
}
func (announceRepo *mockAnnounceRepo) Update(id uint, announce domains.Announce) error {
	return announceRepo.mockUpdate(id, announce)
}
func (announceRepo *mockAnnounceRepo) Delete(id uint) error {
	return announceRepo.mockDelete(id)
}

// Fake achieveRepo that implements AchieveRepo interface
type mockAchieveRepo struct {
	mockInitialize             func(*sql.DB)
	mockSelectAll              func(bool) ([]domains.Achieve, error)
	mockSelectAllUnpublished   func() ([]domains.Achieve, error)
	mockSelectById             func(uint) (domains.Achieve, error)
	mockSelectAllGroupedByYear func() ([]domains.AchieveYearGroup, error)
	mockInsert                 func(domains.Achieve) error
	mockUpdate                 func(uint, domains.Achieve) error
	mockPublish                func([]uint) error
	mockDelete                 func(uint) error
}

// Implement methods of AchieveRepo interface with mocked implementations
func (achieveRepo *mockAchieveRepo) Initialize(db *sql.DB) {}
func (achieveRepo *mockAchieveRepo) SelectAll(publishedOnly bool) ([]domains.Achieve, error) {
	return achieveRepo.mockSelectAll(publishedOnly)
}
func (achieveRepo *mockAchieveRepo) SelectAllUnpublished() ([]domains.Achieve, error) {
	return achieveRepo.mockSelectAllUnpublished()
}
func (achieveRepo *mockAchieveRepo) SelectById(id uint) (domains.Achieve, error) {
	return achieveRepo.mockSelectById(id)
}
func (achieveRepo *mockAchieveRepo) SelectAllGroupedByYear() ([]domains.AchieveYearGroup, error) {
	return achieveRepo.mockSelectAllGroupedByYear()
}
func (achieveRepo *mockAchieveRepo) Insert(achieve domains.Achieve) error {
	return achieveRepo.mockInsert(achieve)
}
func (achieveRepo *mockAchieveRepo) Update(id uint, achieve domains.Achieve) error {
	return achieveRepo.mockUpdate(id, achieve)
}
func (achieveRepo *mockAchieveRepo) Publish(ids []uint) error {
	return achieveRepo.mockPublish(ids)
}
func (achieveRepo *mockAchieveRepo) Delete(id uint) error {
	return achieveRepo.mockDelete(id)
}

// Fake semesterRepo that implements SemesterRepo interface
type mockSemesterRepo struct {
	mockInitialize           func(*sql.DB)
	mockSelectAll            func(bool) ([]domains.Semester, error)
	mockSelectAllUnpublished func() ([]domains.Semester, error)
	mockSelectBySemesterId   func(string) (domains.Semester, error)
	mockInsert               func(domains.Semester) error
	mockUpdate               func(string, domains.Semester) error
	mockPublish              func([]string) error
	mockDelete               func(string) error
}

// Implement methods of SemesterRepo interface with mocked implementations
func (semesterRepo *mockSemesterRepo) Initialize(db *sql.DB) {}
func (semesterRepo *mockSemesterRepo) SelectAll(publishedOnly bool) ([]domains.Semester, error) {
	return semesterRepo.mockSelectAll(publishedOnly)
}
func (semesterRepo *mockSemesterRepo) SelectAllUnpublished() ([]domains.Semester, error) {
	return semesterRepo.mockSelectAllUnpublished()
}
func (semesterRepo *mockSemesterRepo) SelectBySemesterId(semesterId string) (domains.Semester, error) {
	return semesterRepo.mockSelectBySemesterId(semesterId)
}
func (semesterRepo *mockSemesterRepo) Insert(semester domains.Semester) error {
	return semesterRepo.mockInsert(semester)
}
func (semesterRepo *mockSemesterRepo) Update(semesterId string, semester domains.Semester) error {
	return semesterRepo.mockUpdate(semesterId, semester)
}
func (semesterRepo *mockSemesterRepo) Publish(semesterIds []string) error {
	return semesterRepo.mockPublish(semesterIds)
}
func (semesterRepo *mockSemesterRepo) Delete(semesterId string) error {
	return semesterRepo.mockDelete(semesterId)
}

// Fake sessionRepo that implements SessionRepo interface
type mockSessionRepo struct {
	mockInitialize           func(*sql.DB)
	mockSelectAllByClassId   func(string, bool) ([]domains.Session, error)
	mockSelectAllUnpublished func() ([]domains.Session, error)
	mockSelectBySessionId    func(uint) (domains.Session, error)
	mockInsert               func([]domains.Session) error
	mockUpdate               func(uint, domains.Session) error
	mockPublish              func([]uint) error
	mockDelete               func([]uint) error
}

// Implement methods of SessionRepo interface with mocked implementations
func (sessionRepo *mockSessionRepo) Initialize(db *sql.DB) {}
func (sessionRepo *mockSessionRepo) SelectAllByClassId(classId string, publishedOnly bool) ([]domains.Session, error) {
	return sessionRepo.mockSelectAllByClassId(classId, publishedOnly)
}
func (sessionRepo *mockSessionRepo) SelectAllUnpublished() ([]domains.Session, error) {
	return sessionRepo.mockSelectAllUnpublished()
}
func (sessionRepo *mockSessionRepo) SelectBySessionId(id uint) (domains.Session, error) {
	return sessionRepo.mockSelectBySessionId(id)
}
func (sessionRepo *mockSessionRepo) Insert(sessions []domains.Session) error {
	return sessionRepo.mockInsert(sessions)
}
func (sessionRepo *mockSessionRepo) Update(id uint, session domains.Session) error {
	return sessionRepo.mockUpdate(id, session)
}
func (sessionRepo *mockSessionRepo) Publish(ids []uint) error {
	return sessionRepo.mockPublish(ids)
}
func (sessionRepo *mockSessionRepo) Delete(ids []uint) error {
	return sessionRepo.mockDelete(ids)
}

// Fake userRepo that implements UserRepo interface
type mockUserRepo struct {
	mockInitialize         func(*sql.DB)
	mockSelectAll          func(string, int, int) ([]domains.User, error)
	mockSelectById         func(uint) (domains.User, error)
	mockSelectByGuardianId func(uint) ([]domains.User, error)
	mockInsert             func(domains.User) error
	mockUpdate             func(uint, domains.User) error
	mockDelete             func(uint) error
}

// Implement methods of UserRepo interface with mocked implementations
func (userRepo *mockUserRepo) Initialize(db *sql.DB) {}
func (userRepo *mockUserRepo) SelectAll(search string, pageSize, offset int) ([]domains.User, error) {
	return userRepo.mockSelectAll(search, pageSize, offset)
}
func (userRepo *mockUserRepo) SelectById(id uint) (domains.User, error) {
	return userRepo.mockSelectById(id)
}
func (userRepo *mockUserRepo) SelectByGuardianId(guardianId uint) ([]domains.User, error) {
	return userRepo.mockSelectByGuardianId(guardianId)
}
func (userRepo *mockUserRepo) Insert(user domains.User) error {
	return userRepo.mockInsert(user)
}
func (userRepo *mockUserRepo) Update(id uint, user domains.User) error {
	return userRepo.mockUpdate(id, user)
}
func (userRepo *mockUserRepo) Delete(id uint) error {
	return userRepo.mockDelete(id)
}

type mockFamilyRepo struct {
	mockInitialize           func(*sql.DB)
	mockSelectById           func(uint) (domains.Family, error)
	mockSelectByPrimaryEmail func(string) (domains.Family, error)
	mockInsert               func(domains.Family) error
	mockUpdate               func(uint, domains.Family) error
	mockDelete               func(uint) error
}

// Implement methods of UserRepo interface with mocked implementations
func (familyRepo *mockFamilyRepo) Initialize(db *sql.DB) {}

func (familyRepo *mockFamilyRepo) SelectById(id uint) (domains.Family, error) {
	return familyRepo.mockSelectById(id)
}
func (familyRepo *mockFamilyRepo) SelectByPrimaryEmail(primary_email string) (domains.Family, error) {
	return familyRepo.mockSelectByPrimaryEmail(primary_email)
}
func (familyRepo *mockFamilyRepo) Insert(family domains.Family) error {
	return familyRepo.mockInsert(family)
}
func (familyRepo *mockFamilyRepo) Update(id uint, family domains.Family) error {
	return familyRepo.mockUpdate(id, family)
}
func (familyRepo *mockFamilyRepo) Delete(id uint) error {
	return familyRepo.mockDelete(id)
}
