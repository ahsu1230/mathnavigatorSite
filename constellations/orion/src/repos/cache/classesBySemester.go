package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

var KEY_PROGRAM_CLASSES_BY_SEMESTER = "all_programs_semesters_classes"

func GetAllProgramClassesBySemester() ([]domains.ProgramClassesBySemester, error) {
	logger.Debug("Retrieving cache.ProgramClassesBySemester", logger.Fields{
		"key": KEY_PROGRAM_CLASSES_BY_SEMESTER,
	})
	key := KEY_PROGRAM_CLASSES_BY_SEMESTER
	value, err := redis.Values(GetConn().Do("HGETALL", key))
	if err != nil {
		err = appErrors.WrapRedisGet(err, key)
		return []domains.ProgramClassesBySemester{}, err
	}
	
	var list []domains.ProgramClassesBySemester
	err = redis.ScanStruct(value, &list)
	if err != nil {
		err = appErrors.WrapRedisScan(err, key, value)
		return []domains.ProgramClassesBySemester{}, err
	}
	return list, nil
}

func SetAllProgramClassesBySemester(list []domains.ProgramClassesBySemester) error {
	logger.Debug("Setting cache.ProgramClassesBySemester", logger.Fields{
		"key": KEY_PROGRAM_CLASSES_BY_SEMESTER,
		"list": list,
	})
	key := KEY_PROGRAM_CLASSES_BY_SEMESTER
	_, err := GetConn().Do("HMSET", redis.Args{key}.AddFlat(list))
	if err != nil {
		return appErrors.WrapRedisSet(err, key, list)
	}
	return nil
}