package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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
	body := strings.NewReader(`{
		"programId": "ap_calculus",
		"name": "AP Calculus",
		"grade1": 9,
		"grade2": 12,
		"description": "Students should take this course if they aim to take the AP Calculus Exam"
	}`)
	sendPostRequest(hostAddress+"/api/programs/create", body)

	body1 := strings.NewReader(`{
				"programId": "ap_java",
				"name": "AP Java",
				"grade1": 10,
				"grade2": 12,
				"description": "Students should take this course if they aim to take the AP Java Exam"
			}`)
	sendPostRequest(hostAddress+"/api/programs/create", body1)

	body2 := strings.NewReader(`{
				"programId": "sat_math",
				"name": "SAT Math",
				"grade1": 8,
				"grade2": 11,
				"description": "Students should take the course if they aim to take the SAT Math Exam"
			}`)
	sendPostRequest(hostAddress+"/api/programs/create", body2)

	// Create semesters
	body3 := strings.NewReader(`{
				"semesterId": "2020_fall",
				"title": "Fall 2020"
			}`)
	sendPostRequest(hostAddress+"/api/semesters/create", body3)

	body4 := strings.NewReader(`{
				"semesterId": "2021_summer",
				"title": "Summer 2021"
			}`)
	sendPostRequest(hostAddress+"/api/semesters/create", body4)

	body5 := strings.NewReader(`{
				"semesterId": "2021_winter",
				"title": "Winter 2021"
			}`)
	sendPostRequest(hostAddress+"/api/semesters/create", body5)

	// Create locations
	body6 := strings.NewReader(`{
				"locationId": "xkcd",
				"street": "11300 Gainsborough Rd",
				"city": "Potomac",
				"state": "MD",
				"zipcode": "20854"
			}`)
	sendPostRequest(hostAddress+"/api/locations/create", body6)

	body7 := strings.NewReader(`{
				"locationId": "house1",
				"street": "123 Sesame St",
				"city": "Rockville",
				"state": "MD",
				"zipcode": "20854"
			}`)
	sendPostRequest(hostAddress+"/api/locations/create", body7)

	// Create achievements
	body8 := strings.NewReader(`{
				"year": 2020,
				"message": "100% of students scored above a 1550 on SAT!"
			}`)
	sendPostRequest(hostAddress+"/api/achievements/create", body8)

	body9 := strings.NewReader(`{
				"year": 2020,
				"message": "5 students scored an 800 on SAT Math!"
			}`)
	sendPostRequest(hostAddress+"/api/achievements/create", body9)

	body10 := strings.NewReader(`{
				"year": 2019,
				"message": "10 students scored a 5 on AP Java!"
			}`)
	sendPostRequest(hostAddress+"/api/achievements/create", body10)

	// Create announcements

	body11 := strings.NewReader(`{
								var now = time.Now().UTC()
								"postedAt": now,
								"author": "Author Name",
								"message": "The Summer 2020 session of SAT Math will start on Tuesday, June 30.",
								"onHomePage": true
							}`)
	sendPostRequest(hostAddress+"/api/announcements/create", body11)

	body12 := strings.NewReader(`{
								"postedAt": 2020-06-30 00:04:24.034802 +0000 UTC,
								"author": "Harry Potter",
								"message": "The Fall 2021 session of AP Java will start on Tuesday, September 30.",
								"onHomePage": false
							}`)
	sendPostRequest(hostAddress+"/api/announcements/create", body12)

	// Create classes
	body13 := strings.NewReader(`{
			var now = time.Now().UTC()
			var later1 = now.Add(time.Hour * 24 * 30)
			"programId": "ap_calculus",
			"semesterId": "2020_fall",
			"classKey": "class1"
			"classId": "ap_calculus_2020_fall_class1",
			"locationId": "wchs",
			"times": "1 pm - 2 pm",
			"startDate": now,
			"endDate": later1
		}`)
	sendPostRequest(hostAddress+"/api/classes/create", body13)

	body14 := strings.NewReader(`{
			var now = time.Now().UTC()
			var later1 = now.Add(time.Hour * 24 * 30)
			"programId": "ap_java",
			"semesterId": "2020_fall",
			"classKey": "class1"
			"classId": "ap_java_2020_fall_class1",
			"locationId": "house1",
			"times": "1 pm - 2 pm",
			"startDate": now,
			"endDate": later1
		}`)
	sendPostRequest(hostAddress+"/api/classes/create", body14)

	// Create sessions
	body15 := strings.NewReader(`{
		now := time.Now().UTC()
		"classId": "id_1",
		"startsAt": now,
		"endsAt": now,
		"cancelled": false
	}`)
	sendPostRequest(hostAddress+"/api/sessions/create", body15)
}

func sendPostRequest(url string, body io.Reader) {
	resp, err := http.Post(url, "application/json; charset=UTF-8", body)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Fatalln("Response status was not successful.", resp)
	}
}
