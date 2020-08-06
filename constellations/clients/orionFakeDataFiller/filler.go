package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// This CLI sends http requests to a LOCAL orion webserver.
// This script should not be used for a non-local orion webserver!
// Make sure orion is healthy before running this CLI
//
// You can run this CLI using:
// go run filler.go
// go run filler.go --orionHost http://localhost:8080
//
// or via a binary:
// go build filler.go
// ./filler

var DEFAULT_LOCAL_ORION_HOST string = "http://localhost:8001"

func main() {
	log.Println("Orion Fake Data Filler Client starting...")

	var orionHost string
	flag.StringVar(&orionHost, "orionHost", DEFAULT_LOCAL_ORION_HOST, "Host address of an Orion webserver")
	flag.Parse()

	runFiller(orionHost)

	log.Println("Done filling orion")
}

func runFiller(hostAddress string) {
	// Create programs
	createProgram(
		hostAddress,
		"ap_calculus",
		"AP Calculus",
		9,
		12,
		"Students should take this course if they aim to take the AP Calculus Exam",
	)

	createProgram(
		hostAddress,
		"ap_java",
		"AP Java",
		10,
		12,
		"Students should take this course if they aim to take the AP Java Exam",
	)

	createProgram(
		hostAddress,
		"sat_math",
		"SAT Math",
		8,
		11,
		"Students should take the course if they aim to take the SAT Math Exam",
	)

	// Create semesters
	createSemester(
		hostAddress,
		"2020_fall",
		"Fall 2020",
	)

	createSemester(
		hostAddress,
		"2021_summer",
		"Summer 2021",
	)

	createSemester(
		hostAddress,
		"2021_winter",
		"Winter 2021",
	)

	// Create locations
	createLocation(
		hostAddress,
		"wchs",
		"11300 Gainsborough Rd",
		"Potomac",
		"MD",
		"20854",
	)

	createLocation(
		hostAddress,
		"house1",
		"123 Sesame St",
		"Rockville",
		"MD",
		"20854",
	)

	// Create achievements
	createAchieve(
		hostAddress,
		"2020",
		"100% of students scored above a 1550 on SAT!",
	)

	createAchieve(
		hostAddress,
		"2020",
		"5 students scored an 800 on SAT Math!",
	)

	createAchieve(
		hostAddress,
		"2019",
		"10 students scored a 5 on AP Java!",
	)

	// Create announcements
	createAnnounce(
		hostAddress,
		"Author Name",
		"The Summer 2020 session of SAT Math",
		"true",
	)

	createAnnounce(
		hostAddress,
		"Harry Potter",
		"The Summer 2021 session of SAT Math",
		"false",
	)

	// Create classes
	createClass(
		hostAddress,
		"ap_calculus",
		"2020_fall",
		"class1",
		"ap_calculus_2020_fall_class1",
		"wchs",
		"Tues 1pm - 2pm",
	)

	createClass(
		hostAddress,
		"ap_java",
		"2020_fall",
		"class1",
		"ap_java_2020_fall_class1",
		"house1",
		"Tues 1pm - 2pm",
	)

	// Create sessions
	createSessions(
		hostAddress,
		"ap_java_2020_fall_class1",
		"false",
		3,
	)

	createSessions(
		hostAddress,
		"ap_calculus_2020_fall_class1",
		"true",
		2,
	)

	// Create Accounts
	createAccount(
		hostAddress,
		"emailaddress1@gmail.com",
		"jhdgjhddjhdjuj",
	)
	createAccount(hostAddress,
		"emailaddress2@gmail.com",
		"2redssssa",
	)

	// Create Users
	createUser(
		hostAddress,
		"Joe",
		"Smith",
		"",
		"emailaddress1@gmail.com",
		"301-123-4567",
		0,
		1,
		"notes1",
		"schoolone",
		2001,
	)

	createUser(
		hostAddress,
		"Billy",
		"Bob",
		"Joe",
		"emailaddress2@gmail.com",
		"301-123-4568",
		1,
		2,
		"notes2",
		"schooltwo",
		2002,
	)

	// Create transactions
	createTransaction(
		hostAddress,
		100,
		"pay_paypal",
		"notes1",
		1,
	)

	createTransaction(
		hostAddress,
		101,
		"pay_cash",
		"notes2",
		2,
	)
}

func createProgram(hostAddress, programId string, name string, grade1, grade2 int, description string) error {
	programBody := strings.NewReader(fmt.Sprintf(`{
		"programId": "%s",
		"name": "%s",
		"grade1": %d,
		"grade2": %d,
		"description": "%s"
	}`, programId, name, grade1, grade2, description))
	log.Println("Creating program " + programId + "...")
	sendPostRequest(hostAddress+"/api/programs/create", programBody)
	return nil
}

func createSemester(hostAddress, semesterId string, title string) error {
	semesterBody := strings.NewReader(fmt.Sprintf(`{
		"semesterId": "%s",
		"title": "%s"
	}`, semesterId, title))
	log.Println("Creating semester " + semesterId + "...")
	sendPostRequest(hostAddress+"/api/semesters/create", semesterBody)
	return nil
}

func createLocation(hostAddress, locationId, street, city, state, zipcode string) error {
	locationBody := strings.NewReader(fmt.Sprintf(`{
		"locationId": "%s",
		"street": "%s",
		"city": "%s",
		"state": "%s",
		"zipcode": "%s"
	}`, locationId, street, city, state, zipcode))
	log.Println("Creating location " + locationId + "...")
	sendPostRequest(hostAddress+"/api/locations/create", locationBody)
	return nil
}

func createAchieve(hostAddress, year, message string) error {
	achieveBody := strings.NewReader(fmt.Sprintf(`{
		"year": %s,
		"message": "%s"
	}`, year, message))
	log.Println("Creating achievement " + message + "...")
	sendPostRequest(hostAddress+"/api/achievements/create", achieveBody)
	return nil
}

func createAnnounce(hostAddress, author, message, onHomePage string) error {
	now := time.Now().UTC()
	nowJson, _ := now.MarshalJSON()

	announceBody := strings.NewReader(fmt.Sprintf(`{
		"postedAt": %s,
		"author": "%s",
		"message": "%s",
		"onHomePage": %s
	}`, nowJson, author, message, onHomePage))
	log.Println("Creating announcement " + message + "...")
	sendPostRequest(hostAddress+"/api/announcements/create", announceBody)
	return nil
}

func createClass(hostAddress, programId, semesterId, classKey, classId, locationId, times string) error {
	now := time.Now().UTC()
	nowJson, _ := now.MarshalJSON()
	var later = now.Add(time.Hour * 24 * 30)
	laterJson, _ := later.MarshalJSON()

	classBody := strings.NewReader(fmt.Sprintf(`{
		"programId": "%s",
		"semesterId": "%s",
		"classKey": "%s",
		"classId": "%s",
		"locationId": "%s",
		"times": "%s",
		"startDate": %s,
		"endDate": %s
	}`, programId, semesterId, classKey, classId, locationId, times, nowJson, laterJson))
	log.Println("Creating class " + classId + "...")
	sendPostRequest(hostAddress+"/api/classes/create", classBody)
	return nil
}

func createSessions(hostAddress, classId, cancelled string, numSessions int) error {
	// Create session takes in a list
	var body = "["
	now := time.Now().UTC()
	start := now

	for i := 0; i < numSessions; i++ {
		startJson, _ := start.MarshalJSON()
		end := start.Add(time.Hour * 2)
		endJson, _ := end.MarshalJSON()

		body += fmt.Sprintf(`{
			"classId": "%s",
			"startsAt": %s,
			"endsAt": %s,
			"cancelled": %s
		}`, classId, startJson, endJson, cancelled)

		if i < numSessions-1 {
			body += ","
		}

		start = start.Add(time.Hour * 24 * 7)
	}
	body += "]"
	sessionBody := strings.NewReader(body)
	log.Println("Creating session for " + classId + "...")
	sendPostRequest(hostAddress+"/api/sessions/create", sessionBody)

	return nil
}

func createUser(hostAddress, first_name, last_name, middle_name, email, phone string, is_guardian int, account_id int, notes, school string, graduation_year int) error {
	userBody := strings.NewReader(fmt.Sprintf(`{
		"firstName": "%s",
		"lastName": "%s",
		"middleName": "%s",
		"email": "%s",
		"phone": "%s",
		"isGuardian": %d,
		"accountId": %d,
		"notes": "%s",
		"school": "%s",
		"graduationYear": %d
	}`, first_name, last_name, middle_name, email, phone, is_guardian, account_id, notes, school, graduation_year))
	log.Println("Creating user " + first_name + "...")
	sendPostRequest(hostAddress+"/api/users/create", userBody)
	return nil
}

func createAccount(hostAddress, primary_email, password string) error {
	accountBody := strings.NewReader(fmt.Sprintf(`{
		"primaryEmail": "%s",
		"password": "%s"
	}`, primary_email, password))
	log.Println("Creating account " + primary_email + "...")
	sendPostRequest(hostAddress+"/api/accounts/create", accountBody)
	return nil
}

func createTransaction(hostAddress string, amount int, paymentType string, paymentNotes string, accountId int) error {
	transactionBody := strings.NewReader(fmt.Sprintf(`{
		"amount": %d,
		"paymentType": "%s",
		"paymentNotes": "%s",
		"accountId": %d
		}`, amount, paymentType, paymentNotes, accountId))
	log.Println("Creating transaction " + paymentNotes + "...")
	sendPostRequest(hostAddress+"/api/transactions/create", transactionBody)
	return nil
}

func sendPostRequest(url string, body io.Reader) {
	resp, err := http.Post(url, "application/json; charset=UTF-8", body)
	if err != nil {
		log.Println("Post request was not fulfilled.", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Println("Response status was not successful.", resp)
	}
}
