package semesters

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/orion/database"
)

func GetAllSemesters() []Semester {
	var semesterList []Semester
	database.DbSqlx.Select(&semesterList, "SELECT * FROM semesters")
	return semesterList
}

func GetSemesterById(semesterId string) (Semester, error) {
	semester := Semester{}
	sqlStatement := "SELECT * FROM semesters WHERE semester_id=?"
	err := database.DbSqlx.Get(&semester, sqlStatement, semesterId)
	return semester, err
}

func InsertSemester(semester Semester) error {
	semesterId := semester.SemesterId
	now := utils.TimestampNow()
	db := database.DbSqlx
	sqlStatement := "INSERT INTO semesters " +
		"(created_at, updated_at, deleted_at, semester_id, title) " +
		"VALUES (:createdAt, :updatedAt, :deletedAt, :semesterId, :title)"
	parameters := map[string]interface{}{
		"createdAt":  now,
		"updatedAt":  now,
		"deletedAt":  nil,
		"semesterId": semesterId,
		"title":      semester.Title,
	}
	_, err := db.NamedExec(sqlStatement, parameters)
	return err
}

func UpdateSemesterById(oldSemesterId string, semester Semester) error {
	now := utils.TimestampNow()
	db := database.DbSqlx

	sqlStatement := "UPDATE semesters SET " +
		"updated_at=:updatedAt, " +
		"name=:name, " +
		"semester_id=:semesterId, " +
		"title=:title, " +
		"WHERE semester_id=:oldSemesterId"
	parameters := map[string]interface{}{
		"updatedAt":     now,
		"semesterId":    semester.SemesterId,
		"title":         semester.Title,
		"oldSemesterId": oldSemesterId,
	}
	_, err := db.NamedExec(sqlStatement, parameters)
	return err
}

func DeleteSemesterById(semesterId string) error {
	sqlStatement := "DELETE FROM semesters WHERE semester_id=:semesterId"
	parameters := map[string]interface{}{
		"semesterId": semesterId,
	}
	_, err := database.DbSqlx.NamedExec(sqlStatement, parameters)
	return err
}
