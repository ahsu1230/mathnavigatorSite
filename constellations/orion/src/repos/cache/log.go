package cache

import (
	"context"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
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

func logWithContext(ctx context.Context, label, key string) {
	requestUuid := ctx.Value(domains.REQUEST_UUID)
	logger.Debug(label, logger.Fields{
		"requestUuid": requestUuid,
		"key":         key,
	})
}
