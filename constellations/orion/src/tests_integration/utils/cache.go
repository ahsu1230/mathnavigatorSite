package utils

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/cache"
)

func SetupTestCache(host string, port int, password string) {
	logger.Message("Establishing Cache connection")
	cache.Init(host, port, password)
}

func CloseCache() {
	logger.Message("Closing Cache connection")
	cache.Close()
}
