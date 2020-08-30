package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/clients/orionFakeFiller/account1"
	"github.com/ahsu1230/mathnavigatorSite/constellations/clients/orionFakeFiller/account2"
	"github.com/ahsu1230/mathnavigatorSite/constellations/clients/orionFakeFiller/utils"
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

	utils.InitHostAddress(orionHost)
	runFiller()

	log.Println("Done filling orion")
}

func runFiller() {
	// Create programs
	createProgram(
		"ap_calculus",
		"AP Calculus",
		9,
		12,
		"Students should take this course if they aim to take the AP Calculus Exam",
	)

	createProgram(
		"ap_java",
		"AP Java",
		10,
		12,
		"Students should take this course if they aim to take the AP Java Exam",
	)

	createProgram(
		"sat_math",
		"SAT Math",
		8,
		11,
		"Students should take the course if they aim to take the SAT Math Exam",
	)

	createProgram(
		"amc_prep",
		"AMC Prep",
		9,
		12,
		"Students should take the course if they aim to take the AMC Test",
	)

	// Create semesters
	createSemester(
		"2020_fall",
		"Fall 2020",
	)

	createSemester(
		"2021_summer",
		"Summer 2021",
	)

	createSemester(
		"2021_winter",
		"Winter 2021",
	)

	// Create locations
	createLocation(
		"wchs",
		"11300 Gainsborough Rd",
		"Potomac",
		"MD",
		"20854",
	)

	createLocation(
		"house1",
		"123 Sesame St",
		"Rockville",
		"MD",
		"20854",
	)

	// Create achievements
	createAchieve(
		"2020",
		"100% of students scored above a 1550 on SAT!",
	)

	createAchieve(
		"2020",
		"5 students scored an 800 on SAT Math!",
	)

	createAchieve(
		"2019",
		"10 students scored a 5 on AP Java!",
	)

	// Create announcements
	createAnnounce(
		"Albus Dumbledore",
		"The Summer 2020 session of SAT Math will begin soon. Enrollments are now open!",
		"true",
	)

	createAnnounce(
		"Harry Potter",
		"The Summer 2021 SAT1 Math will temporarily moved to another room. Please check the class page for more information.",
		"false",
	)

	// Create classes
	createClass(
		"ap_calculus",
		"2020_fall",
		"class1",
		"ap_calculus_2020_fall_class1",
		"wchs",
		"Tues 1pm - 2pm",
	)

	createClass(
		"ap_java",
		"2020_fall",
		"class1",
		"ap_java_2020_fall_class1",
		"house1",
		"Wed 1pm - 2pm",
	)

	createClass(
		"amc_prep",
		"2020_fall",
		"class1",
		"amc_prep_2020_fall_class1",
		"house1",
		"Thurs 1pm - 2pm",
	)

	createClass(
		"ap_calculus",
		"2021_summer",
		"class1",
		"ap_calculus_2021_summer_class1",
		"wchs",
		"Tues 5pm - 7pm",
	)

	createClass(
		"ap_calculus",
		"2021_summer",
		"class2",
		"ap_calculus_2021_summer_class2",
		"wchs",
		"Tues 1pm - 2pm",
	)

	createClass(
		"ap_calculus",
		"2021_summer",
		"class3",
		"ap_calculus_2021_summer_class3",
		"wchs",
		"Wed 1pm - 2pm",
	)

	createClass(
		"ap_java",
		"2021_summer",
		"class1",
		"ap_java_2021_summer_class1",
		"house1",
		"Tues 5pm - 7pm",
	)

	createClass(
		"ap_calculus",
		"2021_winter",
		"class1",
		"ap_calculus_2021_winter_class1",
		"wchs",
		"Tues 1pm - 2pm",
	)

	createClass(
		"sat_math",
		"2021_winter",
		"class1",
		"sat_math_2021_winter_class1",
		"wchs",
		"Tues 1pm - 2pm",
	)

	// Create sessions
	createSessions(
		"ap_java_2020_fall_class1",
		"false",
		3,
	)

	createSessions(
		"ap_calculus_2020_fall_class1",
		"true",
		2,
	)

	account1.Fill()
	account2.Fill()
}

func createProgram(programId string, name string, grade1, grade2 int, description string) error {
	programBody := strings.NewReader(fmt.Sprintf(`{
		"programId": "%s",
		"name": "%s",
		"grade1": %d,
		"grade2": %d,
		"description": "%s"
	}`, programId, name, grade1, grade2, description))
	log.Println("Creating program " + programId + "...")
	utils.SendPostRequest("/api/programs/create", programBody)
	return nil
}

func createSemester(semesterId string, title string) error {
	semesterBody := strings.NewReader(fmt.Sprintf(`{
		"semesterId": "%s",
		"title": "%s"
	}`, semesterId, title))
	log.Println("Creating semester " + semesterId + "...")
	utils.SendPostRequest("/api/semesters/create", semesterBody)
	return nil
}

func createLocation(locationId, street, city, state, zipcode string) error {
	locationBody := strings.NewReader(fmt.Sprintf(`{
		"locationId": "%s",
		"street": "%s",
		"city": "%s",
		"state": "%s",
		"zipcode": "%s"
	}`, locationId, street, city, state, zipcode))
	log.Println("Creating location " + locationId + "...")
	utils.SendPostRequest("/api/locations/create", locationBody)
	return nil
}

func createAchieve(year, message string) error {
	achieveBody := strings.NewReader(fmt.Sprintf(`{
		"year": %s,
		"message": "%s"
	}`, year, message))
	log.Println("Creating achievement " + message + "...")
	utils.SendPostRequest("/api/achievements/create", achieveBody)
	return nil
}

func createAnnounce(author, message, onHomePage string) error {
	now := time.Now().UTC()
	nowJson, _ := now.MarshalJSON()

	announceBody := strings.NewReader(fmt.Sprintf(`{
		"postedAt": %s,
		"author": "%s",
		"message": "%s",
		"onHomePage": %s
	}`, nowJson, author, message, onHomePage))
	log.Println("Creating announcement " + message + "...")
	utils.SendPostRequest("/api/announcements/create", announceBody)
	return nil
}

func createClass(programId, semesterId, classKey, classId, locationId, times string) error {
	now := time.Now().UTC()
	nowJson, _ := now.MarshalJSON()
	var later = now.Add(time.Hour * 24 * 30)
	laterJson, _ := later.MarshalJSON()
	priceLump := 800

	classBody := strings.NewReader(fmt.Sprintf(`{
		"programId": "%s",
		"semesterId": "%s",
		"classKey": "%s",
		"classId": "%s",
		"locationId": "%s",
		"times": "%s",
		"startDate": %s,
		"endDate": %s,
		"priceLump": %d
	}`,
		programId,
		semesterId,
		classKey,
		classId,
		locationId,
		times,
		nowJson,
		laterJson,
		priceLump,
	))
	log.Println("Creating class " + classId + "...")
	utils.SendPostRequest("/api/classes/create", classBody)
	return nil
}

func createSessions(classId, cancelled string, numSessions int) error {
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
	utils.SendPostRequest("/api/sessions/create", sessionBody)
	return nil
}
