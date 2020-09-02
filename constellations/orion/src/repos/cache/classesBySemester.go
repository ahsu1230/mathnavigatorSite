package cache

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/gomodule/redigo/redis"
)

var KEY_PROGRAM_CLASSES_BY_SEMESTER = "all_programs_semesters_classes"

func GetAllProgramClassesBySemester() ([]domains.ProgramClassesBySemester, error) {
	key := KEY_PROGRAM_CLASSES_BY_SEMESTER
	logger.Debug("Retrieving cache.ProgramClassesBySemester", logger.Fields{
		"key": key,
	})

	conn, err := getConn()
	if err != nil {
		return []domains.ProgramClassesBySemester{}, err
	}
	defer conn.Close()

	listStr, err := redis.String(conn.Do("GET", key))
	if err != nil {
		err = appErrors.WrapRedisGet(err, key)
		return []domains.ProgramClassesBySemester{}, err
	}

	b := []byte(listStr)
	var list []domains.ProgramClassesBySemester
	err = json.Unmarshal(b, &list)
	if err != nil {
		err = appErrors.WrapMarshalJSON("Error (%v) unmarshaling from redis", err)
		return []domains.ProgramClassesBySemester{}, err
	}

	return list, nil
}

func SetAllProgramClassesBySemester(list []domains.ProgramClassesBySemester) error {
	key := KEY_PROGRAM_CLASSES_BY_SEMESTER
	logger.Debug("Setting cache.ProgramClassesBySemester", logger.Fields{
		"key":  key,
		"list": list,
	})

	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	b, err := json.Marshal(&list)
	if err != nil {
		return appErrors.WrapMarshalJSON("Error (%v) marshaling JSON to redis", err)
	}

	_, err = conn.Do("SET", key, string(b))
	if err != nil {
		return appErrors.WrapRedisSet(err, key, list)
	}
	return nil
}
