package cache

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"time"
)

var KEY_PROGRAM_CLASSES_BY_SEMESTER = "all_programs_semesters_classes"
var DURATION_KEY_PROGRAM_CLASSES_BY_SEMESTER = time.Minute * 10

func GetAllProgramClassesBySemester() ([]domains.ProgramClassesBySemester, error) {
	key := KEY_PROGRAM_CLASSES_BY_SEMESTER
	logger.Debug("Retrieving cache.ProgramClassesBySemester", logger.Fields{
		"key": key,
	})

	if CacheDb == nil {
		err := appErrors.ERR_REDIS_UNAVAILABLE
		return []domains.ProgramClassesBySemester{}, err
	}

	listStr, err := CacheDb.Get(ctx, key).Result()
	if err != nil {
		logCacheMiss(key, err)
		return []domains.ProgramClassesBySemester{}, err
	} else {
		logCacheHit(key)
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

	if CacheDb == nil {
		return appErrors.ERR_REDIS_UNAVAILABLE
	}

	// Marshal list into JSON
	b, err := json.Marshal(&list)
	if err != nil {
		return appErrors.WrapMarshalJSON("Error (%v) marshaling JSON to redis", err)
	}

	// Store into cache as string value
	err = CacheDb.Set(
		ctx,
		key,
		string(b),
		DURATION_KEY_PROGRAM_CLASSES_BY_SEMESTER,
	).Err()
	if err != nil {
		return appErrors.WrapRedisSet(err, key, list)
	}

	return nil
}
