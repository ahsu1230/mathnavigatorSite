package achieve

// import (
// 	"github.com/ahsu1230/mathnavigatorSite/orion/controllers/utils"
// 	"github.com/ahsu1230/mathnavigatorSite/orion/database"
// )

// func GetAllAchievements() []Achieve {
// 	var achieveList []Achieve
// 	database.DbSqlx.Select(&achieveList, "SELECT * FROM achievements")
// 	return achieveList
// }

// func GetAchievementById(id uint) (Achieve, error) {
// 	achieve := Achieve{}
// 	sqlStatement := "SELECT * FROM achievements WHERE id=?"
// 	err := database.DbSqlx.Get(&achieve, sqlStatement, id)
// 	return achieve, err
// }

// func InsertAchievement(achieve Achieve) error {
// 	now := utils.TimestampNow()
// 	db := database.DbSqlx
// 	sqlStatement := "INSERT INTO achievements " +
// 		"(created_at, updated_at, deleted_at, year, message) " +
// 		"VALUES (:createdAt, :updatedAt, :deletedAt, :year, :message)"
// 	parameters := map[string]interface{}{
// 		"createdAt": now,
// 		"updatedAt": now,
// 		"deletedAt": nil,
// 		"year":      achieve.Year,
// 		"message":   achieve.Message,
// 	}
// 	_, err := db.NamedExec(sqlStatement, parameters)
// 	return err
// }

// func UpdateAchievementById(id uint, achieve Achieve) error {
// 	now := utils.TimestampNow()
// 	db := database.DbSqlx

// 	sqlStatement := "UPDATE achievements SET " +
// 		"updated_at=:updatedAt, " +
// 		"year=:year, " +
// 		"message=:message, " +
// 		"WHERE id=:id"
// 	parameters := map[string]interface{}{
// 		"updatedAt": now,
// 		"year":      achieve.Year,
// 		"message":   achieve.Message,
// 	}
// 	_, err := db.NamedExec(sqlStatement, parameters)
// 	return err
// }

// func DeleteAchievementById(id uint) error {
// 	sqlStatement := "DELETE FROM achievements WHERE id=:id"
// 	parameters := map[string]interface{}{
// 		"id": id,
// 	}
// 	_, err := database.DbSqlx.NamedExec(sqlStatement, parameters)
// 	return err
// }
