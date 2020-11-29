package testUtils

import (
	"context"
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
var AccountRepo mockAccountRepo
var UserRepo mockUserRepo
var UserClassRepo mockUserClassRepo
var AskForHelpRepo mockAskForHelpRepo
var TransactionRepo mockTransactionRepo
var UserAfhRepo mockUserAfhRepo

// Fake programRepo that implements ProgramRepo interface
type mockProgramRepo struct {
	MockInitialize        func(context.Context, *sql.DB)
	MockSelectAll         func(context.Context) ([]domains.Program, error)
	MockSelectByProgramId func(context.Context, string) (domains.Program, error)
	MockInsert            func(context.Context, domains.Program) (uint, error)
	MockUpdate            func(context.Context, string, domains.Program) error
	MockDelete            func(context.Context, string) error
}

// Implement methods of ProgramRepo interface with mocked implementations
func (programRepo *mockProgramRepo) Initialize(ctx context.Context, db *sql.DB) {}
func (programRepo *mockProgramRepo) SelectAll(ctx context.Context) ([]domains.Program, error) {
	return programRepo.MockSelectAll(ctx)
}
func (programRepo *mockProgramRepo) SelectByProgramId(ctx context.Context, programId string) (domains.Program, error) {
	return programRepo.MockSelectByProgramId(ctx, programId)
}
func (programRepo *mockProgramRepo) Insert(ctx context.Context, program domains.Program) (uint, error) {
	return programRepo.MockInsert(ctx, program)
}
func (programRepo *mockProgramRepo) Update(ctx context.Context, programId string, program domains.Program) error {
	return programRepo.MockUpdate(ctx, programId, program)
}
func (programRepo *mockProgramRepo) Delete(ctx context.Context, programId string) error {
	return programRepo.MockDelete(ctx, programId)
}

// Fake classRepo that implements ClassRepo interface
type mockClassRepo struct {
	MockInitialize                   func(context.Context, *sql.DB)
	MockSelectAll                    func(context.Context, bool) ([]domains.Class, error)
	MockSelectAllUnpublished         func(context.Context) ([]domains.Class, error)
	MockSelectByClassId              func(context.Context, string) (domains.Class, error)
	MockSelectByProgramId            func(context.Context, string) ([]domains.Class, error)
	MockSelectBySemesterId           func(context.Context, string) ([]domains.Class, error)
	MockSelectByProgramAndSemesterId func(context.Context, string, string) ([]domains.Class, error)
	MockInsert                       func(context.Context, domains.Class) (uint, error)
	MockUpdate                       func(context.Context, string, domains.Class) error
	MockPublish                      func(context.Context, []string) []error
	MockDelete                       func(context.Context, string) error
	MockArchive                      func(context.Context, string) error
}

// Implement methods of ClassRepo interface with mocked implementations
func (classRepo *mockClassRepo) Initialize(ctx context.Context, db *sql.DB) {}
func (classRepo *mockClassRepo) SelectAll(ctx context.Context, publishedOnly bool) ([]domains.Class, error) {
	return classRepo.MockSelectAll(ctx, publishedOnly)
}
func (classRepo *mockClassRepo) SelectAllUnpublished(ctx context.Context) ([]domains.Class, error) {
	return classRepo.MockSelectAllUnpublished(ctx)
}
func (classRepo *mockClassRepo) SelectByClassId(ctx context.Context, classId string) (domains.Class, error) {
	return classRepo.MockSelectByClassId(ctx, classId)
}
func (classRepo *mockClassRepo) SelectByProgramId(ctx context.Context, programId string) ([]domains.Class, error) {
	return classRepo.MockSelectByProgramId(ctx, programId)
}
func (classRepo *mockClassRepo) SelectBySemesterId(ctx context.Context, semesterId string) ([]domains.Class, error) {
	return classRepo.MockSelectBySemesterId(ctx, semesterId)
}
func (classRepo *mockClassRepo) SelectByProgramAndSemesterId(ctx context.Context, programId, semesterId string) ([]domains.Class, error) {
	return classRepo.MockSelectByProgramAndSemesterId(ctx, programId, semesterId)
}
func (classRepo *mockClassRepo) Insert(ctx context.Context, class domains.Class) (uint, error) {
	return classRepo.MockInsert(ctx, class)
}
func (classRepo *mockClassRepo) Update(ctx context.Context, classId string, class domains.Class) error {
	return classRepo.MockUpdate(ctx, classId, class)
}
func (classRepo *mockClassRepo) Publish(ctx context.Context, classIds []string) []error {
	return classRepo.MockPublish(ctx, classIds)
}
func (classRepo *mockClassRepo) Delete(ctx context.Context, classId string) error {
	return classRepo.MockDelete(ctx, classId)
}
func (classRepo *mockClassRepo) Archive(ctx context.Context, classId string) error {
	return classRepo.MockArchive(ctx, classId)
}

// Fake locationRepo that implements LocationRepo interface
type mockLocationRepo struct {
	MockInitialize         func(context.Context, *sql.DB)
	MockSelectAll          func(context.Context) ([]domains.Location, error)
	MockSelectByLocationId func(context.Context, string) (domains.Location, error)
	MockInsert             func(context.Context, domains.Location) (uint, error)
	MockUpdate             func(context.Context, string, domains.Location) error
	MockDelete             func(context.Context, string) error
}

// Implement methods of LocationRepo interface with mocked implementations
func (locationRepo *mockLocationRepo) Initialize(ctx context.Context, db *sql.DB) {}
func (locationRepo *mockLocationRepo) SelectAll(ctx context.Context) ([]domains.Location, error) {
	return locationRepo.MockSelectAll(ctx)
}
func (locationRepo *mockLocationRepo) SelectByLocationId(ctx context.Context, locationId string) (domains.Location, error) {
	return locationRepo.MockSelectByLocationId(ctx, locationId)
}
func (locationRepo *mockLocationRepo) Insert(ctx context.Context, location domains.Location) (uint, error) {
	return locationRepo.MockInsert(ctx, location)
}
func (locationRepo *mockLocationRepo) Update(ctx context.Context, locationId string, location domains.Location) error {
	return locationRepo.MockUpdate(ctx, locationId, location)
}
func (locationRepo *mockLocationRepo) Delete(ctx context.Context, locationId string) error {
	return locationRepo.MockDelete(ctx, locationId)
}

// Fake announceRepo that implements AnnounceRepo interface
type mockAnnounceRepo struct {
	MockInitialize         func(context.Context, *sql.DB)
	MockSelectAll          func(context.Context) ([]domains.Announce, error)
	MockSelectByAnnounceId func(context.Context, uint) (domains.Announce, error)
	MockInsert             func(context.Context, domains.Announce) (uint, error)
	MockUpdate             func(context.Context, uint, domains.Announce) error
	MockDelete             func(context.Context, uint) error
}

// Implement methods of AnnounceRepo interface with mocked implementations
func (announceRepo *mockAnnounceRepo) Initialize(ctx context.Context, db *sql.DB) {}
func (announceRepo *mockAnnounceRepo) SelectAll(ctx context.Context) ([]domains.Announce, error) {
	return announceRepo.MockSelectAll(ctx)
}
func (announceRepo *mockAnnounceRepo) SelectByAnnounceId(ctx context.Context, id uint) (domains.Announce, error) {
	return announceRepo.MockSelectByAnnounceId(ctx, id)
}
func (announceRepo *mockAnnounceRepo) Insert(ctx context.Context, announce domains.Announce) (uint, error) {
	return announceRepo.MockInsert(ctx, announce)
}
func (announceRepo *mockAnnounceRepo) Update(ctx context.Context, id uint, announce domains.Announce) error {
	return announceRepo.MockUpdate(ctx, id, announce)
}
func (announceRepo *mockAnnounceRepo) Delete(ctx context.Context, id uint) error {
	return announceRepo.MockDelete(ctx, id)
}

// Fake achieveRepo that implements AchieveRepo interface
type mockAchieveRepo struct {
	MockInitialize             func(context.Context, *sql.DB)
	MockSelectAll              func(context.Context) ([]domains.Achieve, error)
	MockSelectById             func(context.Context, uint) (domains.Achieve, error)
	MockSelectAllGroupedByYear func(context.Context) ([]domains.AchieveYearGroup, error)
	MockInsert                 func(context.Context, domains.Achieve) (uint, error)
	MockUpdate                 func(context.Context, uint, domains.Achieve) error
	MockDelete                 func(context.Context, uint) error
	MockArchive                func(context.Context, string) error
}

// Implement methods of AchieveRepo interface with mocked implementations
func (achieveRepo *mockAchieveRepo) Initialize(ctx context.Context, db *sql.DB) {}
func (achieveRepo *mockAchieveRepo) SelectAll(ctx context.Context) ([]domains.Achieve, error) {
	return achieveRepo.MockSelectAll(ctx)
}
func (achieveRepo *mockAchieveRepo) SelectById(ctx context.Context, id uint) (domains.Achieve, error) {
	return achieveRepo.MockSelectById(ctx, id)
}
func (achieveRepo *mockAchieveRepo) SelectAllGroupedByYear(ctx context.Context) ([]domains.AchieveYearGroup, error) {
	return achieveRepo.MockSelectAllGroupedByYear(ctx)
}
func (achieveRepo *mockAchieveRepo) Insert(ctx context.Context, achieve domains.Achieve) (uint, error) {
	return achieveRepo.MockInsert(ctx, achieve)
}
func (achieveRepo *mockAchieveRepo) Update(ctx context.Context, id uint, achieve domains.Achieve) error {
	return achieveRepo.MockUpdate(ctx, id, achieve)
}
func (achieveRepo *mockAchieveRepo) Delete(ctx context.Context, id uint) error {
	return achieveRepo.MockDelete(ctx, id)
}

// Fake semesterRepo that implements SemesterRepo interface
type mockSemesterRepo struct {
	MockInitialize         func(context.Context, *sql.DB)
	MockSelectAll          func(context.Context) ([]domains.Semester, error)
	MockSelectBySemesterId func(context.Context, string) (domains.Semester, error)
	MockInsert             func(context.Context, domains.Semester) (uint, error)
	MockUpdate             func(context.Context, string, domains.Semester) error
	MockDelete             func(context.Context, string) error
	MockArchive            func(context.Context, string) error
}

// Implement methods of SemesterRepo interface with mocked implementations
func (semesterRepo *mockSemesterRepo) Initialize(ctx context.Context, db *sql.DB) {}
func (semesterRepo *mockSemesterRepo) SelectAll(ctx context.Context) ([]domains.Semester, error) {
	return semesterRepo.MockSelectAll(ctx)
}
func (semesterRepo *mockSemesterRepo) SelectBySemesterId(ctx context.Context, semesterId string) (domains.Semester, error) {
	return semesterRepo.MockSelectBySemesterId(ctx, semesterId)
}
func (semesterRepo *mockSemesterRepo) Insert(ctx context.Context, semester domains.Semester) (uint, error) {
	return semesterRepo.MockInsert(ctx, semester)
}
func (semesterRepo *mockSemesterRepo) Update(ctx context.Context, semesterId string, semester domains.Semester) error {
	return semesterRepo.MockUpdate(ctx, semesterId, semester)
}
func (semesterRepo *mockSemesterRepo) Delete(ctx context.Context, semesterId string) error {
	return semesterRepo.MockDelete(ctx, semesterId)
}
func (semesterRepo *mockSemesterRepo) Archive(ctx context.Context, semesterId string) error {
	return semesterRepo.MockArchive(ctx, semesterId)
}

// Fake sessionRepo that implements SessionRepo interface
type mockSessionRepo struct {
	MockInitialize         func(context.Context, *sql.DB)
	MockSelectAllByClassId func(context.Context, string) ([]domains.Session, error)
	MockSelectBySessionId  func(context.Context, uint) (domains.Session, error)
	MockInsert             func(context.Context, []domains.Session) ([]uint, []error)
	MockUpdate             func(context.Context, uint, domains.Session) error
	MockDelete             func(context.Context, []uint) []error
}

// Implement methods of SessionRepo interface with mocked implementations
func (sessionRepo *mockSessionRepo) Initialize(ctx context.Context, db *sql.DB) {}
func (sessionRepo *mockSessionRepo) SelectAllByClassId(ctx context.Context, classId string) ([]domains.Session, error) {
	return sessionRepo.MockSelectAllByClassId(ctx, classId)
}
func (sessionRepo *mockSessionRepo) SelectBySessionId(ctx context.Context, id uint) (domains.Session, error) {
	return sessionRepo.MockSelectBySessionId(ctx, id)
}
func (sessionRepo *mockSessionRepo) Insert(ctx context.Context, sessions []domains.Session) ([]uint, []error) {
	return sessionRepo.MockInsert(ctx, sessions)
}
func (sessionRepo *mockSessionRepo) Update(ctx context.Context, id uint, session domains.Session) error {
	return sessionRepo.MockUpdate(ctx, id, session)
}
func (sessionRepo *mockSessionRepo) Delete(ctx context.Context, ids []uint) []error {
	return sessionRepo.MockDelete(ctx, ids)
}

// Fake userRepo that implements UserRepo interface
type mockUserRepo struct {
	MockInitialize        func(context.Context, *sql.DB)
	MockSearchUsers       func(context.Context, string) ([]domains.User, error)
	MockSelectAll         func(context.Context, string, int, int) ([]domains.User, error)
	MockSelectById        func(context.Context, uint) (domains.User, error)
	MockSelectByAccountId func(context.Context, uint) ([]domains.User, error)
	MockSelectByEmail     func(context.Context, string) (domains.User, error)
	MockSelectByNew       func(context.Context) ([]domains.User, error)
	MockInsert            func(context.Context, domains.User) (uint, error)
	MockUpdate            func(context.Context, uint, domains.User) error
	MockDelete            func(context.Context, uint) error
	MockFullDelete        func(context.Context, uint) error
}

// Implement methods of UserRepo interface with mocked implementations
func (userRepo *mockUserRepo) Initialize(ctx context.Context, db *sql.DB) {}

func (userRepo *mockUserRepo) SearchUsers(ctx context.Context, search string) ([]domains.User, error) {
	return userRepo.MockSearchUsers(ctx, search)
}

func (userRepo *mockUserRepo) SelectAll(ctx context.Context, search string, pageSize, offset int) ([]domains.User, error) {
	return userRepo.MockSelectAll(ctx, search, pageSize, offset)
}
func (userRepo *mockUserRepo) SelectById(ctx context.Context, id uint) (domains.User, error) {
	return userRepo.MockSelectById(ctx, id)
}
func (userRepo *mockUserRepo) SelectByAccountId(ctx context.Context, accountId uint) ([]domains.User, error) {
	return userRepo.MockSelectByAccountId(ctx, accountId)
}
func (userRepo *mockUserRepo) SelectByEmail(ctx context.Context, email string) (domains.User, error) {
	return userRepo.MockSelectByEmail(ctx, email)
}
func (userRepo *mockUserRepo) SelectByNew(ctx context.Context) ([]domains.User, error) {
	return userRepo.MockSelectByNew(ctx)
}
func (userRepo *mockUserRepo) Insert(ctx context.Context, user domains.User) (uint, error) {
	return userRepo.MockInsert(ctx, user)
}
func (userRepo *mockUserRepo) Update(ctx context.Context, id uint, user domains.User) error {
	return userRepo.MockUpdate(ctx, id, user)
}
func (userRepo *mockUserRepo) Delete(ctx context.Context, id uint) error {
	return userRepo.MockDelete(ctx, id)
}
func (userRepo *mockUserRepo) FullDelete(ctx context.Context, id uint) error {
	return userRepo.MockFullDelete(ctx, id)
}

// Fake userRepo that implements UserRepo interface
type mockUserClassRepo struct {
	MockInitialize           func(context.Context, *sql.DB)
	MockSelectByClassId      func(context.Context, string) ([]domains.UserClass, error)
	MockSelectByUserId       func(context.Context, uint) ([]domains.UserClass, error)
	MockSelectByUserAndClass func(context.Context, uint, string) (domains.UserClass, error)
	MockSelectByNew          func(context.Context) ([]domains.UserClass, error)
	MockInsert               func(context.Context, domains.UserClass) (uint, error)
	MockUpdate               func(context.Context, uint, domains.UserClass) error
	MockDelete               func(context.Context, uint) error
}

// Implement methods of UserRepo interface with mocked implementations
func (userClassesRepo *mockUserClassRepo) Initialize(ctx context.Context, db *sql.DB) {}

func (userClassesRepo *mockUserClassRepo) SelectByClassId(ctx context.Context, classId string) ([]domains.UserClass, error) {
	return userClassesRepo.MockSelectByClassId(ctx, classId)
}
func (userClassesRepo *mockUserClassRepo) SelectByUserId(ctx context.Context, id uint) ([]domains.UserClass, error) {
	return userClassesRepo.MockSelectByUserId(ctx, id)
}
func (userClassesRepo *mockUserClassRepo) SelectByUserAndClass(ctx context.Context, id uint, classId string) (domains.UserClass, error) {
	return userClassesRepo.MockSelectByUserAndClass(ctx, id, classId)
}
func (userClassesRepo *mockUserClassRepo) SelectByNew(ctx context.Context) ([]domains.UserClass, error) {
	return userClassesRepo.MockSelectByNew(ctx)
}
func (userClassesRepo *mockUserClassRepo) Insert(ctx context.Context, userClasses domains.UserClass) (uint, error) {
	return userClassesRepo.MockInsert(ctx, userClasses)
}
func (userClassesRepo *mockUserClassRepo) Update(ctx context.Context, id uint, userClasses domains.UserClass) error {
	return userClassesRepo.MockUpdate(ctx, id, userClasses)
}
func (userClassesRepo *mockUserClassRepo) Delete(ctx context.Context, id uint) error {
	return userClassesRepo.MockDelete(ctx, id)
}

type mockAccountRepo struct {
	MockInitialize                func(context.Context, *sql.DB)
	MockSelectById                func(context.Context, uint) (domains.Account, error)
	MockSelectByPrimaryEmail      func(context.Context, string) (domains.Account, error)
	MockSelectAllNegativeBalances func(context.Context) ([]domains.AccountBalance, error)
	MockInsertWithUser            func(context.Context, domains.Account, domains.User) (uint, error)
	MockUpdate                    func(context.Context, uint, domains.Account) error
	MockDelete                    func(context.Context, uint) error
	MockFullDelete                func(context.Context, uint) error
}

// Implement methods of UserRepo interface with mocked implementations
func (accountRepo *mockAccountRepo) Initialize(ctx context.Context, db *sql.DB) {}

func (accountRepo *mockAccountRepo) SelectById(ctx context.Context, id uint) (domains.Account, error) {
	return accountRepo.MockSelectById(ctx, id)
}
func (accountRepo *mockAccountRepo) SelectByPrimaryEmail(ctx context.Context, primaryEmail string) (domains.Account, error) {
	return accountRepo.MockSelectByPrimaryEmail(ctx, primaryEmail)
}
func (accountRepo *mockAccountRepo) SelectAllNegativeBalances(ctx context.Context) ([]domains.AccountBalance, error) {
	return accountRepo.MockSelectAllNegativeBalances(ctx)
}
func (accountRepo *mockAccountRepo) InsertWithUser(ctx context.Context, account domains.Account, user domains.User) (uint, error) {
	return accountRepo.MockInsertWithUser(ctx, account, user)
}
func (accountRepo *mockAccountRepo) Update(ctx context.Context, id uint, account domains.Account) error {
	return accountRepo.MockUpdate(ctx, id, account)
}
func (accountRepo *mockAccountRepo) Delete(ctx context.Context, id uint) error {
	return accountRepo.MockDelete(ctx, id)
}
func (accountRepo *mockAccountRepo) FullDelete(ctx context.Context, id uint) error {
	return accountRepo.MockFullDelete(ctx, id)
}

type mockAskForHelpRepo struct {
	MockInitialize func(context.Context, *sql.DB)
	MockSelectAll  func(context.Context) ([]domains.AskForHelp, error)
	MockSelectById func(context.Context, uint) (domains.AskForHelp, error)
	MockInsert     func(context.Context, domains.AskForHelp) (uint, error)
	MockUpdate     func(context.Context, uint, domains.AskForHelp) error
	MockDelete     func(context.Context, uint) error
}

// Implement methods of AFHRepo interface with mocked implementations
func (askForHelpRepo *mockAskForHelpRepo) Initialize(ctx context.Context, db *sql.DB) {}

func (askForHelpRepo *mockAskForHelpRepo) SelectAll(ctx context.Context) ([]domains.AskForHelp, error) {
	return askForHelpRepo.MockSelectAll(ctx)
}
func (askForHelpRepo *mockAskForHelpRepo) SelectById(ctx context.Context, id uint) (domains.AskForHelp, error) {
	return askForHelpRepo.MockSelectById(ctx, id)
}
func (askForHelpRepo *mockAskForHelpRepo) Insert(ctx context.Context, askForHelp domains.AskForHelp) (uint, error) {
	return askForHelpRepo.MockInsert(ctx, askForHelp)
}
func (askForHelpRepo *mockAskForHelpRepo) Update(ctx context.Context, id uint, askForHelp domains.AskForHelp) error {
	return askForHelpRepo.MockUpdate(ctx, id, askForHelp)
}
func (askForHelpRepo mockAskForHelpRepo) Delete(ctx context.Context, id uint) error {
	return askForHelpRepo.MockDelete(ctx, id)
}
func (askForHelpRepo mockAskForHelpRepo) Archive(ctx context.Context, id uint) error {
	return askForHelpRepo.MockDelete(ctx, id)
}

type mockTransactionRepo struct {
	MockInitialize        func(context.Context, *sql.DB)
	MockSelectByAccountId func(context.Context, uint) ([]domains.Transaction, error)
	MockSelectById        func(context.Context, uint) (domains.Transaction, error)
	MockInsert            func(context.Context, domains.Transaction) (uint, error)
	MockUpdate            func(context.Context, uint, domains.Transaction) error
	MockDelete            func(context.Context, uint) error
}

// Implement methods of AFHRepo interface with mocked implementations
func (transactionRepo *mockTransactionRepo) Initialize(ctx context.Context, db *sql.DB) {}

func (transactionRepo *mockTransactionRepo) SelectByAccountId(ctx context.Context, accountId uint) ([]domains.Transaction, error) {
	return transactionRepo.MockSelectByAccountId(ctx, accountId)
}
func (transactionRepo *mockTransactionRepo) SelectById(ctx context.Context, id uint) (domains.Transaction, error) {
	return transactionRepo.MockSelectById(ctx, id)
}
func (transactionRepo *mockTransactionRepo) Insert(ctx context.Context, transaction domains.Transaction) (uint, error) {
	return transactionRepo.MockInsert(ctx, transaction)
}
func (transactionRepo *mockTransactionRepo) Update(ctx context.Context, id uint, transaction domains.Transaction) error {
	return transactionRepo.MockUpdate(ctx, id, transaction)
}
func (transactionRepo mockTransactionRepo) Delete(ctx context.Context, id uint) error {
	return transactionRepo.MockDelete(ctx, id)
}

type mockUserAfhRepo struct {
	MockInitalize       func(context.Context, *sql.DB)
	MockSelectByUserId  func(context.Context, uint) ([]domains.UserAfh, error)
	MockSelectByAfhId   func(context.Context, uint) ([]domains.UserAfh, error)
	MockSelectByBothIds func(context.Context, uint, uint) (domains.UserAfh, error)
	MockSelectByNew     func(context.Context) ([]domains.UserAfh, error)
	MockInsert          func(context.Context, domains.UserAfh) (uint, error)
	MockUpdate          func(context.Context, uint, domains.UserAfh) error
	MockDelete          func(context.Context, uint) error
}

// Implement methods of UserAfhRepo interface with mocked implementations
func (userAfhRepo *mockUserAfhRepo) Initialize(ctx context.Context, db *sql.DB) {}

func (userAfhRepo *mockUserAfhRepo) SelectByUserId(ctx context.Context, userId uint) ([]domains.UserAfh, error) {
	return userAfhRepo.MockSelectByUserId(ctx, userId)
}
func (userAfhRepo *mockUserAfhRepo) SelectByAfhId(ctx context.Context, afhId uint) ([]domains.UserAfh, error) {
	return userAfhRepo.MockSelectByAfhId(ctx, afhId)
}
func (userAfhRepo *mockUserAfhRepo) SelectByBothIds(ctx context.Context, userId, afhId uint) (domains.UserAfh, error) {
	return userAfhRepo.MockSelectByBothIds(ctx, userId, afhId)
}
func (userAfhRepo *mockUserAfhRepo) SelectByNew(ctx context.Context) ([]domains.UserAfh, error) {
	return userAfhRepo.MockSelectByNew(ctx)
}
func (userAfhRepo *mockUserAfhRepo) Insert(ctx context.Context, userAfh domains.UserAfh) (uint, error) {
	return userAfhRepo.MockInsert(ctx, userAfh)
}
func (userAfhRepo *mockUserAfhRepo) Update(ctx context.Context, id uint, userAfh domains.UserAfh) error {
	return userAfhRepo.MockUpdate(ctx, id, userAfh)
}
func (userAfhRepo *mockUserAfhRepo) Delete(ctx context.Context, id uint) error {
	return userAfhRepo.MockDelete(ctx, id)
}
