package testUtils

import (
	"database/sql"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

// Mocked structs
var ProgramRepo mockProgramRepo
var ClassRepo mockClassRepo
var LocationRepo mockLocationRepo
var AnnounceRepo mockAnnounceRepo
var AchieveRepo mockAchieveRepo
var SemesterRepo mockSemesterRepo
var SessionRepo mockSessionRepo
var FamilyRepo mockFamilyRepo
var UserRepo mockUserRepo
var AskForHelpRepo mockAskForHelpRepo
var UserAfhRepo mockUserAfhRepo

// Fake programRepo that implements ProgramRepo interface
type mockProgramRepo struct {
	MockInitialize           func(*sql.DB)
	MockSelectAll            func(bool) ([]domains.Program, error)
	MockSelectAllUnpublished func() ([]domains.Program, error)
	MockSelectByProgramId    func(string) (domains.Program, error)
	MockInsert               func(domains.Program) error
	MockUpdate               func(string, domains.Program) error
	MockPublish              func([]string) error
	MockDelete               func(string) error
}

// Implement methods of ProgramRepo interface with mocked implementations
func (programRepo *mockProgramRepo) Initialize(db *sql.DB) {}
func (programRepo *mockProgramRepo) SelectAll(publishedOnly bool) ([]domains.Program, error) {
	return programRepo.MockSelectAll(publishedOnly)
}
func (programRepo *mockProgramRepo) SelectAllUnpublished() ([]domains.Program, error) {
	return programRepo.MockSelectAllUnpublished()
}
func (programRepo *mockProgramRepo) SelectByProgramId(programId string) (domains.Program, error) {
	return programRepo.MockSelectByProgramId(programId)
}
func (programRepo *mockProgramRepo) Insert(program domains.Program) error {
	return programRepo.MockInsert(program)
}
func (programRepo *mockProgramRepo) Update(programId string, program domains.Program) error {
	return programRepo.MockUpdate(programId, program)
}
func (programRepo *mockProgramRepo) Publish(programIds []string) error {
	return programRepo.MockPublish(programIds)
}
func (programRepo *mockProgramRepo) Delete(programId string) error {
	return programRepo.MockDelete(programId)
}

// Fake classRepo that implements ClassRepo interface
type mockClassRepo struct {
	MockInitialize                   func(*sql.DB)
	MockSelectAll                    func(bool) ([]domains.Class, error)
	MockSelectAllUnpublished         func() ([]domains.Class, error)
	MockSelectByClassId              func(string) (domains.Class, error)
	MockSelectByProgramId            func(string) ([]domains.Class, error)
	MockSelectBySemesterId           func(string) ([]domains.Class, error)
	MockSelectByProgramAndSemesterId func(string, string) ([]domains.Class, error)
	MockInsert                       func(domains.Class) error
	MockUpdate                       func(string, domains.Class) error
	MockPublish                      func([]string) error
	MockDelete                       func(string) error
}

// Implement methods of ClassRepo interface with mocked implementations
func (classRepo *mockClassRepo) Initialize(db *sql.DB) {}
func (classRepo *mockClassRepo) SelectAll(publishedOnly bool) ([]domains.Class, error) {
	return classRepo.MockSelectAll(publishedOnly)
}
func (classRepo *mockClassRepo) SelectAllUnpublished() ([]domains.Class, error) {
	return classRepo.MockSelectAllUnpublished()
}
func (classRepo *mockClassRepo) SelectByClassId(classId string) (domains.Class, error) {
	return classRepo.MockSelectByClassId(classId)
}
func (classRepo *mockClassRepo) SelectByProgramId(programId string) ([]domains.Class, error) {
	return classRepo.MockSelectByProgramId(programId)
}
func (classRepo *mockClassRepo) SelectBySemesterId(semesterId string) ([]domains.Class, error) {
	return classRepo.MockSelectBySemesterId(semesterId)
}
func (classRepo *mockClassRepo) SelectByProgramAndSemesterId(programId, semesterId string) ([]domains.Class, error) {
	return classRepo.MockSelectByProgramAndSemesterId(programId, semesterId)
}
func (classRepo *mockClassRepo) Insert(class domains.Class) error {
	return classRepo.MockInsert(class)
}
func (classRepo *mockClassRepo) Update(classId string, class domains.Class) error {
	return classRepo.MockUpdate(classId, class)
}
func (classRepo *mockClassRepo) Publish(classIds []string) error {
	return classRepo.MockPublish(classIds)
}
func (classRepo *mockClassRepo) Delete(classId string) error {
	return classRepo.MockDelete(classId)
}

// Fake locationRepo that implements LocationRepo interface
type mockLocationRepo struct {
	MockInitialize           func(*sql.DB)
	MockSelectAll            func(bool) ([]domains.Location, error)
	MockSelectAllUnpublished func() ([]domains.Location, error)
	MockSelectByLocationId   func(string) (domains.Location, error)
	MockInsert               func(domains.Location) error
	MockUpdate               func(string, domains.Location) error
	MockPublish              func([]string) error
	MockDelete               func(string) error
}

// Implement methods of LocationRepo interface with mocked implementations
func (locationRepo *mockLocationRepo) Initialize(db *sql.DB) {}
func (locationRepo *mockLocationRepo) SelectAll(publishedOnly bool) ([]domains.Location, error) {
	return locationRepo.MockSelectAll(publishedOnly)
}
func (locationRepo *mockLocationRepo) SelectAllUnpublished() ([]domains.Location, error) {
	return locationRepo.MockSelectAllUnpublished()
}
func (locationRepo *mockLocationRepo) SelectByLocationId(locationId string) (domains.Location, error) {
	return locationRepo.MockSelectByLocationId(locationId)
}
func (locationRepo *mockLocationRepo) Insert(location domains.Location) error {
	return locationRepo.MockInsert(location)
}
func (locationRepo *mockLocationRepo) Update(locationId string, location domains.Location) error {
	return locationRepo.MockUpdate(locationId, location)
}
func (locationRepo *mockLocationRepo) Publish(locationIds []string) error {
	return locationRepo.MockPublish(locationIds)
}
func (locationRepo *mockLocationRepo) Delete(locationId string) error {
	return locationRepo.MockDelete(locationId)
}

// Fake announceRepo that implements AnnounceRepo interface
type mockAnnounceRepo struct {
	MockInitialize         func(*sql.DB)
	MockSelectAll          func() ([]domains.Announce, error)
	MockSelectByAnnounceId func(uint) (domains.Announce, error)
	MockInsert             func(domains.Announce) error
	MockUpdate             func(uint, domains.Announce) error
	MockDelete             func(uint) error
}

// Implement methods of AnnounceRepo interface with mocked implementations
func (announceRepo *mockAnnounceRepo) Initialize(db *sql.DB) {}
func (announceRepo *mockAnnounceRepo) SelectAll() ([]domains.Announce, error) {
	return announceRepo.MockSelectAll()
}
func (announceRepo *mockAnnounceRepo) SelectByAnnounceId(id uint) (domains.Announce, error) {
	return announceRepo.MockSelectByAnnounceId(id)
}
func (announceRepo *mockAnnounceRepo) Insert(announce domains.Announce) error {
	return announceRepo.MockInsert(announce)
}
func (announceRepo *mockAnnounceRepo) Update(id uint, announce domains.Announce) error {
	return announceRepo.MockUpdate(id, announce)
}
func (announceRepo *mockAnnounceRepo) Delete(id uint) error {
	return announceRepo.MockDelete(id)
}

// Fake achieveRepo that implements AchieveRepo interface
type mockAchieveRepo struct {
	MockInitialize             func(*sql.DB)
	MockSelectAll              func(bool) ([]domains.Achieve, error)
	MockSelectAllUnpublished   func() ([]domains.Achieve, error)
	MockSelectById             func(uint) (domains.Achieve, error)
	MockSelectAllGroupedByYear func() ([]domains.AchieveYearGroup, error)
	MockInsert                 func(domains.Achieve) error
	MockUpdate                 func(uint, domains.Achieve) error
	MockPublish                func([]uint) error
	MockDelete                 func(uint) error
}

// Implement methods of AchieveRepo interface with mocked implementations
func (achieveRepo *mockAchieveRepo) Initialize(db *sql.DB) {}
func (achieveRepo *mockAchieveRepo) SelectAll(publishedOnly bool) ([]domains.Achieve, error) {
	return achieveRepo.MockSelectAll(publishedOnly)
}
func (achieveRepo *mockAchieveRepo) SelectAllUnpublished() ([]domains.Achieve, error) {
	return achieveRepo.MockSelectAllUnpublished()
}
func (achieveRepo *mockAchieveRepo) SelectById(id uint) (domains.Achieve, error) {
	return achieveRepo.MockSelectById(id)
}
func (achieveRepo *mockAchieveRepo) SelectAllGroupedByYear() ([]domains.AchieveYearGroup, error) {
	return achieveRepo.MockSelectAllGroupedByYear()
}
func (achieveRepo *mockAchieveRepo) Insert(achieve domains.Achieve) error {
	return achieveRepo.MockInsert(achieve)
}
func (achieveRepo *mockAchieveRepo) Update(id uint, achieve domains.Achieve) error {
	return achieveRepo.MockUpdate(id, achieve)
}
func (achieveRepo *mockAchieveRepo) Publish(ids []uint) error {
	return achieveRepo.MockPublish(ids)
}
func (achieveRepo *mockAchieveRepo) Delete(id uint) error {
	return achieveRepo.MockDelete(id)
}

// Fake semesterRepo that implements SemesterRepo interface
type mockSemesterRepo struct {
	MockInitialize           func(*sql.DB)
	MockSelectAll            func(bool) ([]domains.Semester, error)
	MockSelectAllUnpublished func() ([]domains.Semester, error)
	MockSelectBySemesterId   func(string) (domains.Semester, error)
	MockInsert               func(domains.Semester) error
	MockUpdate               func(string, domains.Semester) error
	MockPublish              func([]string) error
	MockDelete               func(string) error
}

// Implement methods of SemesterRepo interface with mocked implementations
func (semesterRepo *mockSemesterRepo) Initialize(db *sql.DB) {}
func (semesterRepo *mockSemesterRepo) SelectAll(publishedOnly bool) ([]domains.Semester, error) {
	return semesterRepo.MockSelectAll(publishedOnly)
}
func (semesterRepo *mockSemesterRepo) SelectAllUnpublished() ([]domains.Semester, error) {
	return semesterRepo.MockSelectAllUnpublished()
}
func (semesterRepo *mockSemesterRepo) SelectBySemesterId(semesterId string) (domains.Semester, error) {
	return semesterRepo.MockSelectBySemesterId(semesterId)
}
func (semesterRepo *mockSemesterRepo) Insert(semester domains.Semester) error {
	return semesterRepo.MockInsert(semester)
}
func (semesterRepo *mockSemesterRepo) Update(semesterId string, semester domains.Semester) error {
	return semesterRepo.MockUpdate(semesterId, semester)
}
func (semesterRepo *mockSemesterRepo) Publish(semesterIds []string) error {
	return semesterRepo.MockPublish(semesterIds)
}
func (semesterRepo *mockSemesterRepo) Delete(semesterId string) error {
	return semesterRepo.MockDelete(semesterId)
}

// Fake sessionRepo that implements SessionRepo interface
type mockSessionRepo struct {
	MockInitialize           func(*sql.DB)
	MockSelectAllByClassId   func(string, bool) ([]domains.Session, error)
	MockSelectAllUnpublished func() ([]domains.Session, error)
	MockSelectBySessionId    func(uint) (domains.Session, error)
	MockInsert               func([]domains.Session) error
	MockUpdate               func(uint, domains.Session) error
	MockPublish              func([]uint) error
	MockDelete               func([]uint) error
}

// Implement methods of SessionRepo interface with mocked implementations
func (sessionRepo *mockSessionRepo) Initialize(db *sql.DB) {}
func (sessionRepo *mockSessionRepo) SelectAllByClassId(classId string, publishedOnly bool) ([]domains.Session, error) {
	return sessionRepo.MockSelectAllByClassId(classId, publishedOnly)
}
func (sessionRepo *mockSessionRepo) SelectAllUnpublished() ([]domains.Session, error) {
	return sessionRepo.MockSelectAllUnpublished()
}
func (sessionRepo *mockSessionRepo) SelectBySessionId(id uint) (domains.Session, error) {
	return sessionRepo.MockSelectBySessionId(id)
}
func (sessionRepo *mockSessionRepo) Insert(sessions []domains.Session) error {
	return sessionRepo.MockInsert(sessions)
}
func (sessionRepo *mockSessionRepo) Update(id uint, session domains.Session) error {
	return sessionRepo.MockUpdate(id, session)
}
func (sessionRepo *mockSessionRepo) Publish(ids []uint) error {
	return sessionRepo.MockPublish(ids)
}
func (sessionRepo *mockSessionRepo) Delete(ids []uint) error {
	return sessionRepo.MockDelete(ids)
}

// Fake userRepo that implements UserRepo interface
type mockUserRepo struct {
	MockInitialize       func(*sql.DB)
	MockSelectAll        func(string, int, int) ([]domains.User, error)
	MockSelectById       func(uint) (domains.User, error)
	MockSelectByFamilyId func(uint) ([]domains.User, error)
	MockInsert           func(domains.User) error
	MockUpdate           func(uint, domains.User) error
	MockDelete           func(uint) error
}

// Implement methods of UserRepo interface with mocked implementations
func (userRepo *mockUserRepo) Initialize(db *sql.DB) {}
func (userRepo *mockUserRepo) SelectAll(search string, pageSize, offset int) ([]domains.User, error) {
	return userRepo.MockSelectAll(search, pageSize, offset)
}
func (userRepo *mockUserRepo) SelectById(id uint) (domains.User, error) {
	return userRepo.MockSelectById(id)
}
func (userRepo *mockUserRepo) SelectByFamilyId(familyId uint) ([]domains.User, error) {
	return userRepo.MockSelectByFamilyId(familyId)
}
func (userRepo *mockUserRepo) Insert(user domains.User) error {
	return userRepo.MockInsert(user)
}
func (userRepo *mockUserRepo) Update(id uint, user domains.User) error {
	return userRepo.MockUpdate(id, user)
}
func (userRepo *mockUserRepo) Delete(id uint) error {
	return userRepo.MockDelete(id)
}

type mockFamilyRepo struct {
	MockInitialize           func(*sql.DB)
	MockSelectById           func(uint) (domains.Family, error)
	MockSelectByPrimaryEmail func(string) (domains.Family, error)
	MockInsert               func(domains.Family) error
	MockUpdate               func(uint, domains.Family) error
	MockDelete               func(uint) error
}

// Implement methods of UserRepo interface with mocked implementations
func (familyRepo *mockFamilyRepo) Initialize(db *sql.DB) {}

func (familyRepo *mockFamilyRepo) SelectById(id uint) (domains.Family, error) {
	return familyRepo.MockSelectById(id)
}
func (familyRepo *mockFamilyRepo) SelectByPrimaryEmail(primary_email string) (domains.Family, error) {
	return familyRepo.MockSelectByPrimaryEmail(primary_email)
}
func (familyRepo *mockFamilyRepo) Insert(family domains.Family) error {
	return familyRepo.MockInsert(family)
}
func (familyRepo *mockFamilyRepo) Update(id uint, family domains.Family) error {
	return familyRepo.MockUpdate(id, family)
}
func (familyRepo *mockFamilyRepo) Delete(id uint) error {
	return familyRepo.MockDelete(id)
}

type mockAskForHelpRepo struct {
	MockInitialize func(*sql.DB)
	MockSelectAll  func() ([]domains.AskForHelp, error)
	MockSelectById func(uint) (domains.AskForHelp, error)
	MockInsert     func(domains.AskForHelp) error
	MockUpdate     func(uint, domains.AskForHelp) error
	MockDelete     func(uint) error
}

// Implement methods of AFHRepo interface with mocked implementations
func (askForHelpRepo *mockAskForHelpRepo) Initialize(db *sql.DB) {}

func (askForHelpRepo *mockAskForHelpRepo) SelectAll() ([]domains.AskForHelp, error) {
	return askForHelpRepo.MockSelectAll()
}
func (askForHelpRepo *mockAskForHelpRepo) SelectById(id uint) (domains.AskForHelp, error) {
	return askForHelpRepo.MockSelectById(id)
}
func (askForHelpRepo *mockAskForHelpRepo) Insert(askForHelp domains.AskForHelp) error {
	return askForHelpRepo.MockInsert(askForHelp)
}
func (askForHelpRepo *mockAskForHelpRepo) Update(id uint, askForHelp domains.AskForHelp) error {
	return askForHelpRepo.MockUpdate(id, askForHelp)
}
func (askForHelpRepo mockAskForHelpRepo) Delete(id uint) error {
	return askForHelpRepo.MockDelete(id)
}

type mockUserAfhRepo struct {
	MockInitalize       func(*sql.DB)
	MockSelectByUserId  func(uint) ([]domains.UserAfh, error)
	MockSelectByAfhId   func(uint) ([]domains.UserAfh, error)
	MockSelectByBothIds func(uint, uint) (domains.UserAfh, error)
	MockInsert          func(domains.UserAfh) error
	MockUpdate          func(uint, domains.UserAfh) error
	MockDelete          func(uint) error
}

// Implement methods of UserAfhRepo interface with mocked implementations
func (userAfhRepo *mockUserAfhRepo) Initialize(db *sql.DB) {}

func (userAfhRepo *mockUserAfhRepo) SelectByUserId(userId uint) ([]domains.UserAfh, error) {
	return userAfhRepo.MockSelectByUserId(userId)
}
func (userAfhRepo *mockUserAfhRepo) SelectByAfhId(afhId uint) ([]domains.UserAfh, error) {
	return userAfhRepo.MockSelectByAfhId(afhId)
}
func (userAfhRepo *mockUserAfhRepo) SelectByBothIds(userId, afhId uint) (domains.UserAfh, error) {
	return userAfhRepo.MockSelectByBothIds(userId, afhId)
}
func (userAfhRepo *mockUserAfhRepo) Insert(userAfh domains.UserAfh) error {
	return userAfhRepo.MockInsert(userAfh)
}
func (userAfhRepo *mockUserAfhRepo) Update(id uint, userAfh domains.UserAfh) error {
	return userAfhRepo.MockUpdate(id, userAfh)
}
func (userAfhRepo *mockUserAfhRepo) Delete(id uint) error {
	return userAfhRepo.MockDelete(id)
}
