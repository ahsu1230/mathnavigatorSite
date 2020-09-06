package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

func LogControllerMethod(c *gin.Context, label string) {
	v, ok := c.Get(domains.REQUEST_UUID)
	if ok {
		requestUuid := v.(uuid.UUID)
		logger.Debug(label, logger.Fields{
			"requestUuid": requestUuid,
		})
	}
}

func RetrieveContext(c *gin.Context) context.Context {
	return c.Request.Context()
}
