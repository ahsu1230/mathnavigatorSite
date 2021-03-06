package cache

import (
	"context"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/go-redis/redis/v8"
	// "github.com/alicebob/miniredis"
	// "github.com/elliotchance/redismock"
)

var (
	CacheDb *redis.Client
)

func Init(host string, port int, password string) {
	logger.Message("Initializing Redis...")

	connection := fmt.Sprintf("%s:%d", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: password,
		DB:       0, // use default DB
	})

	logger.Message("Pinging Redis...")
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.Error("Error pinging Redis", err, logger.Fields{})
		return
	}

	logger.Message("Redis ready!")
	CacheDb = rdb
}

func InitForMockTest() {
	CacheDb = nil
	// TODO: Use a Redis mock tester
	// mr, err := miniredis.Run()
	// if err != nil {
	// 	panic(err)
	// }
	// client := redis.NewClient(&redis.Options{
	// 	Addr: mr.Addr(),
	// })
	// ctx = context.Background()
	// CacheDb = client
	// return redismock.NewNiceMock(client)
}

func Close() {
	logger.Message("Closing Redis...")
	if CacheDb != nil {
		CacheDb.Close()
		CacheDb = nil
	}
	logger.Message("Redis Closed")
}

// Invalidate method (by key)
func Delete(ctx context.Context, key string) error {
	logWithContext(ctx, "Invalidating cacheKey", key)

	if CacheDb == nil {
		return appErrors.ERR_REDIS_UNAVAILABLE
	}

	err := CacheDb.Del(ctx, key).Err()
	if err != nil {
		logger.Error(
			"Error from DEL redis operation",
			err,
			logger.Fields{"key": key},
		)
		return appErrors.WrapRedisDelete(err, key)
	}

	logWithContext(ctx, "Deleted cacheKey", key)
	return nil
}
