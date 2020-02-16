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

func GetAnnouncementById(announceId string) (Announce, error) {
	var announce Announce
	err := database.DbSqlx.Get(&announce, "SELECT * FROM announces WHERE announce_id=?", announceId)
	return announce, err
}

func InsertAnnouncement(announce Announce) (error) {
	now := utils.TimestampNow()
	sqlStatement := "INSERT INTO announcements " +
		"(created_at, updated_at, deleted_at, announce_id, title, message) " +
        "VALUES (:createdAt, :updatedAt, :deletedAt, :announceId, :title, :message)"
	parameters := map[string]interface{} {
		"createdAt": now,
		"updatedAt": now,
		"deletedAt": nil,
		"announceId": announce.AnnounceId,
		"title": announce.Title,
		"message": announce.Message,
	}
	_, err := database.DbSqlx.NamedExec(sqlStatement, parameters)
	return err
}

func UpdateAnnouncementById(oldAnnounceId string, announce Announce) (error) {
	now := utils.TimestampNow()

	sqlStatement := "UPDATE announcements SET " +
		"updated_at=:updatedAt, " +
		"announce_id=:announceId, " +
		"title=:grade1, " +
		"message=:grade2, " +
		"WHERE announce_id=:oldAnnounceId"
	parameters := map[string]interface{}{
		"updatedAt": now,
		"announceId": announce.AnnounceId,
		"title": announce.Title,
		"message": announce.Message,
		"oldAnnounceId": oldAnnounceId,
	}
	_, err := database.DbSqlx.NamedExec(sqlStatement, parameters)
	return err
}

func DeleteAnnouncementById(announceId string) error {
	_, err := database.DbSqlx.NamedExec("DELETE FROM announces WHERE announce_id=:announceId", map[string]interface{}{"announceId": announceId})
	return err
}