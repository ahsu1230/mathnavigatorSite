package middlewares

import (
	"github.com/gin-contrib/cors"
)

func CreateCorsConfig(config Config) cors.Config {
	configCors := cors.DefaultConfig()
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	configCors.AllowOrigins = []string{config.App.CorsOrigin}
	return configCors
}
