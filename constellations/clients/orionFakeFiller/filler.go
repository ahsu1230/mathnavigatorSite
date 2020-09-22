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
		"wchs",
		"Final Project AMA",
	)
	date = time.Now().Add(time.Hour * 24 * 17)
	afhId2, _ := createAFH(
		date,
		"AP Java Office Hours",
		"programming",
		"wchs",
		"Please bring your exam notes to review!",
	)

	account1.Fill()
	account2.Fill(afhId1, afhId2)
}

func createProgram(programId string, title string, grade1, grade2 int, description string) (uint, error) {
	programBody := strings.NewReader(fmt.Sprintf(`{
		"programId": "%s",
		"title": "%s",
		"grade1": %d,
		"grade2": %d,
		"description": "%s",
		"featured": "none"
	}`, programId, title, grade1, grade2, description))
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

func createClass(programId, semesterId, classKey, classId, locationId, times string) (uint, error) {
	priceLumpSum := 800

	classBody := strings.NewReader(fmt.Sprintf(`{
		"programId": "%s",
		"semesterId": "%s",
		"classKey": "%s",
		"classId": "%s",
		"locationId": "%s",
		"timesStr": "%s",
		"priceLumpSum": %d
	}`,
		programId,
		semesterId,
		classKey,
		classId,
		locationId,
		times,
		priceLumpSum,
	))
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
