package cache

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/gomodule/redigo/redis"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

var (
	// Pool *redis.Pool
	Conn *redis.Conn
)

func Init(host string, port int, password string) {
	logger.Message("Initializing Redis...")

	connection := fmt.Sprintf("%s:%d", host, port)
	conn, err := redis.Dial("tcp", connection)
	if err != nil {
		logger.Error("Failed to connect to Redis", err, logger.Fields{
			"connection": connection,
		})
		return
	}
	Conn = &conn

	if err = GetConn().Err(); err != nil {
		logger.Error("Error connecting to Redis", err, logger.Fields{})
		return
	}

	GetConn().Do("FlushAll")
	logger.Message("Connected to Redis!")
}

func InitForTest() {
	// Conn = redigomock.NewConn()
}

func Close() {
	logger.Message("Closing Redis...")
	if Conn != nil {
		if err := GetConn().Close(); err != nil {
			logger.Error("Error closing Redis", err, logger.Fields{})
			return
		}
		Conn = nil
	}
	logger.Message("Closed")
}

func GetConn() redis.Conn {
	return *Conn
}

// Invalidate method (by key)
func Delete(key string) error {
	logger.Debug("Invalidating cacheKey", logger.Fields{
		"key": key,
	})
	_, err := GetConn().Do("DELETE", redis.Args{key})
	return err
}

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