package importer

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Struct that matches row/json from file
type UserClassFromFile struct {
	classId string `json:"classId"`
	name    string `json:"name"`
	email   string `json:"email"`
}

// User domain object from orion
type UserStudent struct {
	firstName  string `json:"firstName"`
	lastName   string `json:"lastName"`
	email      string `json:"email"`
	isGuardian bool   `json:"isGuardian"`
	// School         string   	`json:"school"`
	// GraduationYear NullUint     `json:"graduationYear"`
	// Notes          NullString   `json:"notes"`
}

// request body object for register endpoint in orion
type RegisterRequest struct {
	student UserStudent `json:"student"`
}

func ImportUserClass(hostAddress string, data []byte) error {
	var rows []UserClassFromFile
	err := json.Unmarshal(data, &rows)
	if err != nil {
		log.Println("Error converting file to JSON")
		return err
	}

	for i := 0; i < len(rows); i++ {
		row := rows[i]
		fullName := strings.Split(row.name, " ")
		firstName := fullName[0]
		lastName := fullName[len(fullName)-1]
		userStudent := UserStudent{
			firstName,
			lastName,
			row.email,
			true,
		}
		req := RegisterRequest{
			student: userStudent,
		}
		classId := row.classId
		log.Println(fmt.Sprintf("Attempting to http POST request for user %s for class %s", row.name, classId))
		url := fmt.Sprintf("%s/api/register/class/%s", hostAddress, classId)
		SendPostRequest(url, req)
	}
	return nil
}
