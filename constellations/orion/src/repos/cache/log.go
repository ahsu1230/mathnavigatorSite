package cache

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func LogCacheHit(key string) {
	logger.Debug("Cache Hit", logger.Fields{
		"key": key,
	})
}

func LogCacheMiss(key string, err error) {
	if errors.Is(err, redis.ErrNil) {
		logger.Debug("Cache Miss", logger.Fields{
			"key": key,
		})
	} else { // Actual error from operation!
		LogError(key, err)
	}
}

func LogError(key string, err error) {
	logger.Error("Error from Cache operation", err, logger.Fields{
		"key": key,
	})
}
