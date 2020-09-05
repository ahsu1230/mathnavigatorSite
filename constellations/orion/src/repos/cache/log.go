package cache

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

func logCacheHit(key string) {
	logger.Debug("Cache Hit", logger.Fields{
		"key": key,
	})
}

func logCacheMiss(key string, err error) {
	if errors.Is(err, redis.Nil) {
		logger.Debug("Cache Miss", logger.Fields{
			"key": key,
		})
	} else { // Actual error from operation!
		logError(key, err)
	}
}

func logError(key string, err error) {
	logger.Error("Error from Cache operation", err, logger.Fields{
		"key": key,
	})
}
