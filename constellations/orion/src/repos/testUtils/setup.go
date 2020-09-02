package testUtils

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/cache"
	"github.com/gomodule/redigo/redis"
)

// Global test variables
var CacheConn redis.Conn

// Utility methods
func init() {
	CacheConn = cache.InitForMockTest()
}
