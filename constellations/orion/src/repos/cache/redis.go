package cache

import (
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"time"
)

var (
	Pool *redis.Pool
)

func Init(host string, port int, password string) {
	logger.Message("Initializing Redis...")

	connection := fmt.Sprintf("%s:%d", host, port)
	Pool = &redis.Pool{
		MaxIdle:         1,
		IdleTimeout:     1 * time.Minute, // time to close an idle connection
		MaxConnLifetime: 1 * time.Hour,   // max time to keep connection
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", connection)
			if err != nil {
				return nil, err
			}

			if _, err := conn.Do("AUTH", password); err != nil {
				logger.Error("Error authenticating to Redis", err, logger.Fields{})
				conn.Close()
				return nil, err
			}

			logger.Message("Initialized & authenticated Redis!")
			conn.Do("FLUSHALL")
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func InitForMockTest() {
	conn := redigomock.NewConn()
	Pool = &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return conn, nil
		},
	}
}

func Close() {
	logger.Message("Closing Redis...")
	if Pool != nil {
		Pool.Close()
		Pool = nil
	}
	logger.Message("Redis Closed")
}

func getConn() (redis.Conn, error) {
	if Pool == nil {
		logger.Debug("Redis Cache is not currently available", logger.Fields{})
		return nil, appErrors.ERR_REDIS_UNAVAILABLE
	}
	return Pool.Get(), nil
}

func FlushAll() {
	logger.Debug("Flush all keys", logger.Fields{})
	conn, err := getConn()
	if err != nil {
		return
	}
	defer conn.Close()

	conn.Do("FLUSHALL")
}

// Invalidate method (by key)
func Delete(key string) error {
	logger.Debug("Invalidating cacheKey", logger.Fields{
		"key": key,
	})

	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("DELETE", redis.Args{key})
	if err != nil {
		return appErrors.WrapRedisDelete(err, key)
	}

	return nil
}
