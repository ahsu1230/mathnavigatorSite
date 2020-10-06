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
	createProgram(`{
		"programId": "ap_calculus",
		"title": "AP Calculus",
		"grade1": 9,
		"grade2": 12,
		"subject": "math",
		"description": "Students should take this course if they aim to take the AP Calculus Exam.",
		"featured": "popular"
	}`, "ap_calculus")

	createProgram(`{
		"programId": "ap_java",
		"title": "AP Java",
		"grade1": 10,
		"grade2": 12,
		"subject": "programming",
		"description": "Students should take this course if they aim to take the AP Java Exam.",
		"featured": "new"
	}`, "ap_java")

	createProgram(`{
		"programId": "sat_math",
		"title": "SAT Math",
		"grade1": 8,
		"grade2": 12,
		"subject": "math",
		"description": "Students should take the course if they aim to take the SAT Math Exam.",
		"featured": "popular"
	}`, "sat_math")

	createProgram(`{
		"programId": "amc_prep",
		"title": "AMC 10 Preparation",
		"grade1": 9,
		"grade2": 10,
		"subject": "math",
		"description": "Students should take the course if they aim to take the AMC Competition Exam.",
		"featured": "none"
	}`, "amc_prep")

	createProgram(`{
		"programId": "sat_writing",
		"title": "SAT Essay Writing",
		"grade1": 9,
		"grade2": 12,
		"subject": "english",
		"description": "Students should take this course to improve fundamental essay writing skills and learn to succeed on the SAT writing section.",
		"featured": "popular"
	}`, "sat_writing")

	// Create semesters
	createSemester(
		"fall",
		2020,
	)

	createSemester(
		"summer",
		2021,
	)

	createSemester(
		"winter",
		2021,
	)

	// Create locations
	createLocation(
		"wchs",
		"Winston Churchill High School",
		"11300 Gainsborough Rd",
		"Potomac",
		"MD",
		"20854",
		"Room 110",
	)

	createLocation(
		"house1",
		"Sesame House",
		"123 Sesame St",
		"Rockville",
		"MD",
		"20854",
		"",
	)

	createOnlineLocation(
		"zoom",
		"Zoom Video Conference",
	)

	// Create achievements
	createAchieve(
		"2020",
		"100% of students scored above a 1550 on SAT!",
	)

	createAchieve(
		"2020",
		"Five students scored an 800 on SAT Math!",
	)

	createAchieve(
		"2020",
		"80% of our SAT2 Subject Math class scored above 750. 50% scored 800!",
	)

	createAchieve(
		"2019",
		"All 12 students scored a 4 or 5 on the AP Java exam.",
	)

	createAchieve(
		"2019",
		"Three students qualified for the National American Mathematics Competition.",
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
	createClass(`{
		"programId": "ap_calculus",
		"semesterId": "2020_fall",
		"classKey": "class1",
		"classId": "ap_calculus_2020_fall_class1",
		"locationId": "wchs",
		"timesStr": "Tues 1pm - 2pm, Fri 3pm - 5pm",
		"priceLumpSum": 800
	}`, "ap_calculus_2020_fall_class1")

	createClass(`{
		"programId": "ap_java",
		"semesterId": "2020_fall",
		"classKey": "class1",
		"classId": "ap_java_2020_fall_class1",
		"locationId": "zoom",
		"timesStr": "Wed 1pm - 2pm",
		"pricePerSession": 60
	}`, "ap_java_2020_fall_class1")

	createClass(`{
		"programId": "amc_prep",
		"semesterId": "2020_fall",
		"classKey": "class1",
		"classId": "amc_prep_2020_fall_class1",
		"locationId": "house1",
		"timesStr": "Thurs 1pm - 2pm",
		"priceLumpSum": 640
	}`, "amc_prep_2020_fall_class1")

	createClass(`{
		"programId": "ap_calculus",
		"semesterId": "2021_summer",
		"classKey": "class1",
		"classId": "ap_calculus_2021_summer_class1",
		"locationId": "wchs",
		"timesStr": "Tues 5pm - 7pm",
		"priceLumpSum": 640
	}`, "ap_calculus_2021_summer_class1")

	createClass(`{
		"programId": "ap_calculus",
		"semesterId": "2021_summer",
		"classKey": "class2",
		"classId": "ap_calculus_2021_summer_class2",
		"locationId": "wchs",
		"timesStr": "Tues 1pm - 2pm",
		"priceLumpSum": 640
	}`, "ap_calculus_2021_summer_class2")

	createClass(`{
		"programId": "ap_calculus",
		"semesterId": "2021_summer",
		"classKey": "class3",
		"classId": "ap_calculus_2021_summer_class3",
		"locationId": "wchs",
		"timesStr": "Wed 1pm - 2pm",
		"priceLumpSum": 640
	}`, "ap_calculus_2021_summer_class3")

	createClass(`{
		"programId": "ap_java",
		"semesterId": "2021_summer",
		"classKey": "class1",
		"classId": "ap_java_2021_summer_class1",
		"locationId": "zoom",
		"timesStr": "Tues 5pm - 7pm",
		"priceLumpSum": 720
	}`, "ap_java_2021_summer_class1")

	createClass(`{
		"programId": "ap_calculus",
		"semesterId": "2021_winter",
		"classKey": "class1",
		"classId": "ap_calculus_2021_winter_class1",
		"locationId": "wchs",
		"timesStr": "Tues 1pm - 2pm",
		"priceLumpSum": 640
	}`, "ap_calculus_2021_winter_class1")

	createClass(`{
		"programId": "sat_math",
		"semesterId": "2021_winter",
		"classKey": "class1",
		"classId": "sat_math_2021_winter_class1",
		"locationId": "zoom",
		"timesStr": "Tues 1pm - 2pm",
		"priceLumpSum": 640
	}`, "sat_math_2021_winter_class1")

	createClass(`{
		"programId": "sat_writing",
		"semesterId": "2021_winter",
		"classKey": "class1",
		"classId": "sat_writing_2021_winter_class1",
		"locationId": "house1",
		"timesStr": "Thurs 1pm - 2pm, Fri 1pm - 2pm",
		"pricePerSession": 80
	}`, "sat_writing_2021_winter_class1")

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

	// Create AFH
	date := time.Now().Add(time.Hour * 24 * 3)
	createAFH(
		date,
		"AP Java Office Hours",
		"programming",
		"wchs",
		"",
	)
	date = time.Now().Add(time.Hour * 24 * 10)
	afhId1, _ := createAFH(
		date,
		"AP Java Office Hours",
		"programming",
		"zoom",
		"Final Project AMA",
	)
	date = time.Now().Add(time.Hour * 24 * 17)
	afhId2, _ := createAFH(
		date,
		"AP Java Office Hours",
		"programming",
		"zoom",
		"Please bring your exam notes to review!",
	)

	date = time.Now().Add(time.Hour * 24 * 6)
	createAFH(
		date,
		"SAT Math Extra Practice",
		"math",
		"wchs",
		"Please come early to give yourself time to prepare!",
	)

	date = time.Now().Add(time.Hour * 24 * 7)
	createAFH(
		date,
		"SAT Math Extra Practice Review",
		"math",
		"wchs",
		"Review from extra practice session.",
	)

	account1.Fill()
	account2.Fill(afhId1, afhId2)
}

func createProgram(body, programId string) (uint, error) {
	programBody := strings.NewReader(body)
	log.Println("Creating program " + programId + "...")
	respBody := utils.SendPostRequest("/api/programs/create", programBody)
	id, _ := utils.GetIdFromBody(respBody)
	return id, nil
}

func createSemester(season string, year int) (uint, error) {
	semesterBody := strings.NewReader(fmt.Sprintf(`{
		"season": "%s",
		"year": %d
	}`, season, year))
	log.Println(fmt.Sprintf("Creating semester %d_%s...", year, season))
	respBody := utils.SendPostRequest("/api/semesters/create", semesterBody)
	id, _ := utils.GetIdFromBody(respBody)
	return id, nil
}

func createLocation(locationId, title, street, city, state, zipcode, room string) (uint, error) {
	locationBody := strings.NewReader(fmt.Sprintf(`{
		"locationId": "%s",
		"title": "%s",
		"street": "%s",
		"city": "%s",
		"state": "%s",
		"zipcode": "%s",
		"room": "%s"
	}`, locationId, title, street, city, state, zipcode, room))
	log.Println("Creating location " + locationId + "...")
	respBody := utils.SendPostRequest("/api/locations/create", locationBody)
	id, _ := utils.GetIdFromBody(respBody)
	return id, nil
}

func createOnlineLocation(locationId, title string) (uint, error) {
	locationBody := strings.NewReader(fmt.Sprintf(`{
		"locationId": "%s",
		"title": "%s",
		"isOnline": true
	}`, locationId, title))
	log.Println("Creating location " + locationId + "...")
	respBody := utils.SendPostRequest("/api/locations/create", locationBody)
	id, _ := utils.GetIdFromBody(respBody)
	return id, nil
}

func createAchieve(year, message string) (uint, error) {
	achieveBody := strings.NewReader(fmt.Sprintf(`{
		"year": %s,
		"message": "%s"
	}`, year, message))
	log.Println("Creating achievement " + message + "...")
	respBody := utils.SendPostRequest("/api/achievements/create", achieveBody)
	id, _ := utils.GetIdFromBody(respBody)
	return id, nil
}

func createAnnounce(author, message, onHomePage string) (uint, error) {
	now := time.Now().UTC()
	nowJson, _ := now.MarshalJSON()

	announceBody := strings.NewReader(fmt.Sprintf(`{
		"postedAt": %s,
		"author": "%s",
		"message": "%s",
		"onHomePage": %s
	}`, nowJson, author, message, onHomePage))
	log.Println("Creating announcement " + message + "...")
	respBody := utils.SendPostRequest("/api/announcements/create", announceBody)
	id, _ := utils.GetIdFromBody(respBody)
	return id, nil
}

func createClass(body, classId string) (uint, error) {
	classBody := strings.NewReader(body)
	log.Println("Creating class " + classId + "...")
	respBody := utils.SendPostRequest("/api/classes/create", classBody)
	id, _ := utils.GetIdFromBody(respBody)
	return id, nil
}

func createSessions(classId, cancelled string, numSessions int) ([]uint, error) {
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
	respBody := utils.SendPostRequest("/api/sessions/create", sessionBody)

	ids, _ := utils.GetIdsFromBody(respBody)
	return ids, nil
}

func createAFH(startsAt time.Time, title, subject, locationId, notes string) (uint, error) {
	startJson, _ := startsAt.MarshalJSON()
	endsAt := startsAt.Add(time.Hour * 1)
	endJson, _ := endsAt.MarshalJSON()
	body := strings.NewReader(fmt.Sprintf(`{
		"startsAt": %s,
		"endsAt": %s,
		"title": "%s",
		"subject": "%s",
		"locationId": "%s",
		"notes": "%s"
	}`, startJson, endJson, title, subject, locationId, notes))
	log.Println("Creating afh " + title + " (" + subject + ") ...")
	respBody := utils.SendPostRequest("/api/askforhelp/create", body)
	id, _ := utils.GetIdFromBody(respBody)
	return id, nil
}
