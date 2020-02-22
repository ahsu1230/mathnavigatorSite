package announce

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/orion/database"
)

func GetAllAnnouncements() []Announce {
	var announceList []Announce
	database.DbSqlx.Select(&announceList, "SELECT * FROM announcements")
	return announceList
}

func GetAnnouncementById(id string) (Announce, error) {
	var announce Announce
	err := database.DbSqlx.Get(&announce, "SELECT * FROM announcements WHERE id=?", id)
	return announce, err
}

func InsertAnnouncement(announce Announce) (error) {
	now := utils.TimestampNow()
	sqlStatement := "INSERT INTO announcements " +
		"(created_at, updated_at, deleted_at, posted_at, author, message) " +
        "VALUES (:createdAt, :updatedAt, :deletedAt, :postedAt, :author, :message)"
	parameters := map[string]interface{} {
		"createdAt": now,
		"updatedAt": now,
		"deletedAt": nil,
		"postedAt": announce.PostedAt,
		"author": announce.Author,
		"message": announce.Message,
	}
	_, err := database.DbSqlx.NamedExec(sqlStatement, parameters)
	return err
}

func UpdateAnnouncementById(id string, announce Announce) (error) {
	now := utils.TimestampNow()

	sqlStatement := "UPDATE announcements SET " +
		"updated_at=:updatedAt, " +
		"posted_at=:postedAt, " +
		"author=:author, " +
		"message=:message " +
		"WHERE id=:id"
	parameters := map[string]interface{}{
		"updatedAt": now,
		"postedAt": announce.PostedAt,
		"author": announce.Author,
		"message": announce.Message,
		"id": id,
	}
	_, err := database.DbSqlx.NamedExec(sqlStatement, parameters)
	return err
}

func DeleteAnnouncementById(id string) error {
	_, err := database.DbSqlx.NamedExec("DELETE FROM announcements WHERE id=:id", map[string]interface{}{"id": id})
	return err
}