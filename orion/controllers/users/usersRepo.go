package users

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/orion/database"
)

func GetAllUsers() []User {
	var userList []User
	database.DbSqlx.Select(&userList, "SELECT * FROM users WHERE deleted_at = null")
	return userList
}

func GetUserById(id uint) (User, error) {
	user := User{}
	err := database.DbSqlx.Get(&user, "SELECT * FROM users WHERE id=? AND deleted_at = null", id)
	return user, err
}

func InsertUser(user User) error {
	now := utils.TimestampNow()
	sqlStatement := "INSERT INTO users " +
		"(created_at, updated_at, deleted_at, first_name, last_name, " +
		"middle_name, email, phone, isGuardian, guardianId) " +
		"VALUES (:createdAt, :updatedAt, :deletedAt, :firstName, :lastName, " +
		":middleName, :email, :phone, :isGuardian, :guardianId)"
	parameters := map[string]interface{}{
		"createdAt":   now,
		"updatedAt":   now,
		"deletedAt":   nil,
		"first_name":  user.FirstName,
		"last_name":   user.LastName,
		"middle_name": user.MiddleName,
		"email":       user.Email,
		"phone":       user.Phone,
		"isGuardian":  user.IsGuardian,
		"guardianId":  user.GuardianId,
	}
	_, err := database.DbSqlx.NamedExec(sqlStatement, parameters)
	return err
}

func UpdateUserById(id uint, user User) error {
	now := utils.TimestampNow()
	sqlStatement := "UPDATE users SET " +
		"updated_at=:updatedAt, " +
		"first_name=:firstName, " +
		"last_name=:lastName, " +
		"middle_name=:middleName, " +
		"email=:email, " +
		"phone=:phone, " +
		"isGuardian=:isGuardian, " +
		"guardianId=:guardianId, " +
		"WHERE id=:id AND deleted_at = null"
	parameters := map[string]interface{}{
		"updatedAt":  now,
		"firstName":  user.FirstName,
		"lastName":   user.LastName,
		"middleName": user.MiddleName,
		"email":      user.Email,
		"phone":      user.Phone,
		"isGuardian": user.IsGuardian,
		"guardianId": user.GuardianId,
	}
	_, err := database.DbSqlx.NamedExec(sqlStatement, parameters)
	return err
}

func DeleteUserById(id uint) error {
	now := utils.TimestampNow()
	sqlStatement := "UPDATE users SET deleted_at=:deletedAt WHERE id=:id"
	parameters := map[string]interface{}{
		"deletedAt": now,
		"id":        id,
	}
	_, err := database.DbSqlx.NamedExec(sqlStatement, parameters)
	return err
}
