package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

func LogControllerMethod(c *gin.Context, label string) (uuid.UUID, bool) {
	v, ok := c.Get("requestUuid")
	if ok {
		requestUuid := v.(uuid.UUID)
		logger.Info(label, logger.Fields{
			"requestUuid": requestUuid,
		})
		return requestUuid, true
	}
	return uuid.UUID{}, false
}
