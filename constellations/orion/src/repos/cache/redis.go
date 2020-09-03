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
	Conn redis.Conn
)

func Init(host string, port int, password string) {
	logger.Message("Initializing Redis...")

	connection := fmt.Sprintf("%s:%d", host, port)
	Pool = &redis.Pool{
		MaxIdle:         3,
		IdleTimeout:     1 * time.Minute, // time to close an idle connection
		MaxConnLifetime: 1 * time.Minute, // max time to keep connection
		Dial: func() (redis.Conn, error) {
			logger.Debug("Retrieving a Redis connection...", logger.Fields{})
			conn, err := redis.Dial("tcp", connection)
			if err != nil {
				logger.Error("Error dialing to Redis", err, logger.Fields{})
				return nil, err
			}

			logger.Debug("Authenticating Redis connection...", logger.Fields{})
			if _, err := conn.Do("AUTH", password); err != nil {
				logger.Error("Error authenticating to Redis", err, logger.Fields{})
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	// Attempt Ping to redis
	logger.Message("Initialized Redis and attempting ping")
	conn, err := getConn()
	if err != nil {
		logger.Error("Error connecting to Redis after initialize", err, logger.Fields{})
		return
	}

	if _, err := redis.String(conn.Do("PING")); err != nil {
		logger.Error("Error pinging Redis", err, logger.Fields{})
		return
	}
	conn.Do("FLUSHALL")
	logger.Message("Redis pinged!")
}

func InitForMockTest() redis.Conn {
	conn := redigomock.NewConn()
	Pool = &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return conn, nil
		},
	}
	return conn
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

// Invalidate method (by key)
func Delete(key string) error {
	logger.Debug("Invalidating cacheKey", logger.Fields{
		"key": key,
	})

	conn, err := getConn()
	if err != nil {
		logger.Debug("Error getting conn when deleting", logger.Fields{})
		return appErrors.WrapRedisUnavailable(err, "Redis unavailable before DELETE")
	}
	defer conn.Close()
	logger.Debug("Retrieved connection", logger.Fields{
		"key": key,
	})

	_, err := conn.Do("DEL", redis.Args{key})
	if err != nil {
		logger.Debug("Error deleting key", logger.Fields{})
		return appErrors.WrapRedisDelete(err, key)
	}

	logger.Debug("Deleted key", logger.Fields{
		"key": key,
	})

	return nil
}
