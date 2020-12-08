package cache

import (
	"context"
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"time"
)

var KEY_PROGRAM_CLASSES_BY_SEMESTER = "all_programs_semesters_classes"
var DURATION_KEY_PROGRAM_CLASSES_BY_SEMESTER = time.Minute * 10

func GetAllProgramClassesBySemester(ctx context.Context) ([]domains.ProgramClassesBySemester, error) {
	key := KEY_PROGRAM_CLASSES_BY_SEMESTER
	logWithContext(ctx, "Retrieving cache.ProgramClassesBySemester", key)

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

func SetAllProgramClassesBySemester(ctx context.Context, list []domains.ProgramClassesBySemester) error {
	key := KEY_PROGRAM_CLASSES_BY_SEMESTER
	logWithContext(ctx, "Setting cache.ProgramClassesBySemester", key)

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
