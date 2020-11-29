package utils

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"net/http/httptest"
	"testing"
	"time"
)

var AccountTonyStark = domains.Account{
	PrimaryEmail: "tony@stark.com",
	Password:     "password1",
}

var AccountNatasha = domains.Account{
	PrimaryEmail: "natasha@shield.com",
	Password:     "password2",
}

var AccountPotter = domains.Account{
	PrimaryEmail: "hpotter@hogwarts.com",
	Password:     "password3",
}

var UserTonyStark = domains.User{
	AccountId:      1,
	FirstName:      "Tony",
	LastName:       "Stark",
	MiddleName:     domains.NewNullString("Edward"),
	Email:          "tony@stark.com",
	Phone:          domains.NewNullString("555-555-0101"),
	IsGuardian:     true,
	Notes:          domains.NewNullString("Avengers CEO"),
	School:         domains.NewNullString(""),
	GraduationYear: domains.NewNullUint(0),
}

var UserMorganStark = domains.User{
	AccountId:      1,
	FirstName:      "Morgan",
	LastName:       "Parker",
	MiddleName:     domains.NewNullString(""),
	Email:          "morgan@stark.com",
	Phone:          domains.NewNullString("555-555-0101"),
	IsGuardian:     false,
	Notes:          domains.NewNullString("Daughter of Tony Stark"),
	School:         domains.NewNullString("Avengers Academy"),
	GraduationYear: domains.NewNullUint(2036),
}

var UserPeterParker = domains.User{
	AccountId:      1,
	FirstName:      "Peter",
	LastName:       "Parker",
	MiddleName:     domains.NewNullString("Benjamin"),
	Email:          "peter@stark.com",
	Phone:          domains.NewNullString("555-555-0102"),
	IsGuardian:     false,
	Notes:          domains.NewNullString("Avengers Intern"),
	School:         domains.NewNullString("Midtown High School"),
	GraduationYear: domains.NewNullUint(2021),
}

var UserNatasha = domains.User{
	AccountId:      2,
	FirstName:      "Natasha",
	LastName:       "Romanova",
	MiddleName:     domains.NewNullString("Alianovna"),
	Email:          "natasha@shield.com",
	Phone:          domains.NewNullString(""),
	IsGuardian:     true,
	Notes:          domains.NewNullString("Secret Agent"),
	School:         domains.NewNullString(""),
	GraduationYear: domains.NewNullUint(0),
}

var UserPotter = domains.User{
	AccountId:      3,
	FirstName:      "Harry",
	LastName:       "Potter",
	MiddleName:     domains.NewNullString("James"),
	Email:          "hpotter@hogwarts.com",
	Phone:          domains.NewNullString("555-555-0103"),
	IsGuardian:     false,
	Notes:          domains.NewNullString("notes123"),
	School:         domains.NewNullString("Hogwarts School"),
	GraduationYear: domains.NewNullUint(2005),
}

func CreateAllTestAccountsAndUsers(t *testing.T) {
	// Create TonyStark's Account with PeterParker as another user
	SendCreateAccountUser(t, true, AccountTonyStark, UserTonyStark) // userId 1, accountId 1
	SendCreateUser(t, true, UserMorganStark)                        // userId 2, accountId 1
	SendCreateUser(t, true, UserPeterParker)                        // userId 3, accountId 1

	// Create Natasha's Account
	SendCreateAccountUser(t, true, AccountNatasha, UserNatasha) // userId 4, accountId 2

	// Create Harry Potter's Account
	SendCreateAccountUser(t, true, AccountPotter, UserPotter) // userId 5, accountId 3
}

func SendCreateLocationWCHS(t *testing.T) (uint, *httptest.ResponseRecorder) {
	return SendCreateLocation(
		t,
		true,
		"wchs",
		"Winston Churchill High School",
		"11300 Gainsborough Road",
		"Potomac",
		"MD",
		"20854",
		"Room 110",
		false,
	)
}

func SendCreateLocationHogwarts(t *testing.T) (uint, *httptest.ResponseRecorder) {
	return SendCreateLocation(
		t,
		true,
		"hogwarts",
		"Hogwarts School of Witchcraft and Wizardry",
		"3950 Dumbledore Ave",
		"Potter",
		"ME",
		"69156",
		"Room 44",
		false,
	)
}

func SendCreateLocationZoom(t *testing.T) (uint, *httptest.ResponseRecorder) {
	return SendCreateLocation(
		t,
		true,
		"wchs",
		"Online Zoom Video Conference",
		"",
		"",
		"",
		"",
		"Link will be provided on Google Classroom",
		true,
	)
}

/*
  Program: "program1"
  Semester: "2020_spring"
  Location: "wchs"
  Class: "program1_2020_spring_classA"
  Afh: 1 (?)
*/
func CreateFullClassAndAfhEnvironment(t *testing.T) {
	time1 := time.Now().UTC()
	time2 := time1.Add(time.Hour * 1)

	SendCreateProgram(
		t,
		true,
		"program1",
		"Program1",
		1,
		3,
		domains.SUBJECT_MATH,
		"description1",
		domains.FEATURED_NONE,
	)
	SendCreateSemester(t, true, domains.SPRING, 2020)
	SendCreateLocationWCHS(t)
	SendCreateClass(
		t,
		true,
		"program1",
		"2020_spring",
		"classA",
		"wchs",
		"2:00pm - 3:00pm",
		0,
		600,
	)
	SendCreateAskForHelp(
		t,
		true,
		time1,
		time2,
		"AP Statistics Help",
		domains.SUBJECT_MATH,
		"wchs",
		"test note 2",
	)
}
