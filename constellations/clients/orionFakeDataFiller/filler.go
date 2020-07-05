package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// This CLI sends http requests to a local orion webserver.
// Make sure orion is healthy before running this CLI
//
// You can run this CLI using:
// go run filler.go
//
// or via a binary:
// go build filler.go
// ./filler

func main() {
	fmt.Println("Orion Fake Data Filler Client starting...")
	runFiller("http://localhost:8001")

	fmt.Println("Done filling orion")
}

func runFiller(hostAddress string) {

	// Create programs
	body := strings.NewReader(createProgram(
		"ap_calculus",
		"AP Calculus",
		"9",
		"12",
		"Students should take this course if they aim to take the AP Calculus Exam",
	))
	sendPostRequest(hostAddress+"/api/programs/create", body)

	body1 := strings.NewReader(createProgram(
		"ap_java",
		"AP Java",
		"10",
		"12",
		"Students should take this course if they aim to take the AP Java Exam",
	))
	sendPostRequest(hostAddress+"/api/programs/create", body1)

	body2 := strings.NewReader(createProgram(
		"sat_math",
		"SAT Math",
		"8",
		"11",
		"Students should take the course if they aim to take the SAT Math Exam",
	))
	sendPostRequest(hostAddress+"/api/programs/create", body2)

	// Create semesters
	body3 := strings.NewReader(createSemester(
		"2020_fall",
		"Fall 2020",
	))
	sendPostRequest(hostAddress+"/api/semesters/create", body3)

	body4 := strings.NewReader(createSemester(
		"2021_summer",
		"Summer 2021",
	))
	sendPostRequest(hostAddress+"/api/semesters/create", body4)

	body5 := strings.NewReader(createSemester(
		"2021_winter",
		"Winter 2021",
	))
	sendPostRequest(hostAddress+"/api/semesters/create", body5)

	// Create locations
	body6 := strings.NewReader(createLocation(
		"wchs",
		"11300 Gainsborough Rd",
		"Potomac",
		"MD",
		"20854",
	))
	sendPostRequest(hostAddress+"/api/locations/create", body6)

	body7 := strings.NewReader(createLocation(
		"house1",
		"123 Sesame St",
		"Rockville",
		"MD",
		"20854",
	))
	sendPostRequest(hostAddress+"/api/locations/create", body7)

	// Create achievements
	body8 := strings.NewReader(createAchieve(
		"2020",
		"100% of students scored above a 1550 on SAT!",
	))
	sendPostRequest(hostAddress+"/api/achievements/create", body8)

	body9 := strings.NewReader(createAchieve(
		"2020",
		"5 students scored an 800 on SAT Math!",
	))
	sendPostRequest(hostAddress+"/api/achievements/create", body9)

	body10 := strings.NewReader(createAchieve(
		"2019",
		"10 students scored a 5 on AP Java!",
	))
	sendPostRequest(hostAddress+"/api/achievements/create", body10)

	// Create announcements
	body11 := strings.NewReader(createAnnounce(
		"Author Name",
		"The Summer 2020 session of SAT Math",
		"true",
	))
	sendPostRequest(hostAddress+"/api/announcements/create", body11)

	body12 := strings.NewReader(createAnnounce(
		"Harry Potter",
		"The Summer 2021 session of SAT Math",
		"false",
	))
	sendPostRequest(hostAddress+"/api/announcements/create", body12)

	// Create classes
	body13 := strings.NewReader(createClass(
		"ap_calculus",
		"2020_fall",
		"class1",
		"ap_calculus_2020_fall_class1",
		"wchs",
		"Tues 1pm - 2pm",
	))
	sendPostRequest(hostAddress+"/api/classes/create", body13)

	body14 := strings.NewReader(createClass(
		"ap_java",
		"2020_fall",
		"class1",
		"ap_java_2020_fall_class1",
		"house1",
		"Tues 1pm - 2pm",
	))
	sendPostRequest(hostAddress+"/api/classes/create", body14)

	// Create sessions
	body15 := strings.NewReader(createSession(
		"id_1",
		"false",
	))
	sendPostRequest(hostAddress+"/api/sessions/create", body15)

	body16 := strings.NewReader(createSession(
		"id_2",
		"true",
	))
	sendPostRequest(hostAddress+"/api/sessions/create", body16)

}

func createProgram(programId string, name string, grade1 string, grade2 string, description string) string {
	programBody := fmt.Sprintf(`{
		"programId": "%s",
		"name": "%s",
		"grade1": %s,
		"grade2": %s,
		"description": "%s"
	}`, programId, name, grade1, grade2, description)
	fmt.Println(programId + " was created")
	return programBody
}

func createSemester(semesterId string, title string) string {
	semesterBody := fmt.Sprintf(`{
		"semesterId": "%s",
		"title": "%s"
	}`, semesterId, title)
	return semesterBody
}

func createLocation(locationId, street, city, state, zipcode string) string {
	locationBody := fmt.Sprintf(`{
		"locationId": "%s",
		"street": "%s",
		"city": "%s",
		"state": "%s",
		"zipcode": "%s"
	}`, locationId, street, city, state, zipcode)
	return locationBody
}

func createAchieve(year, message string) string {
	achieveBody := fmt.Sprintf(`{
		"year": %s,
		"message": "%s"
	}`, year, message)
	return achieveBody
}

func createAnnounce(author, message, onHomePage string) string {
	now := time.Now().UTC()
	nowJson, _ := now.MarshalJSON()

	announceBody := fmt.Sprintf(`{
		"postedAt": %s,
		"author": "%s",
		"message": "%s",
		"onHomePage": %s
	}`, nowJson, author, message, onHomePage)

	return announceBody
}

func createClass(programId, semesterId, classKey, classId, locationId, times string) string {
	now := time.Now().UTC()
	nowJson, _ := now.MarshalJSON()
	var later1 = now.Add(time.Hour * 24 * 30)
	laterJson, _ := later1.MarshalJSON()

	classBody := fmt.Sprintf(`{
		"programId": "%s",
		"semesterId": "%s",
		"classKey": "%s",
		"classId": "%s",
		"locationId": "%s",
		"times": "%s",
		"startDate": %s,
		"endDate": %s
	}`, programId, semesterId, classKey, classId, locationId, times, nowJson, laterJson)
	return classBody
}

func createSession(classId, cancelled string) string {
	now := time.Now().UTC()
	nowJson, _ := now.MarshalJSON()

	sessionBody := fmt.Sprintf(`{
		"classId": "%s"
		"startsAt": %s
		"endsAt": %s
		"cancelled": %s
	}`, classId, nowJson, nowJson, cancelled)
	return sessionBody
}

func sendPostRequest(url string, body io.Reader) {
	resp, err := http.Post(url, "application/json; charset=UTF-8", body)
	if err != nil {
		log.Println("smth failed")
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Println("Response status was not successful.")
	}
}
