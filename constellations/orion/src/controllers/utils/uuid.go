package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

func LogControllerMethod(c *gin.Context, label string) uuid.UUID {
	requestUuid := c.MustGet("requestUuid").(uuid.UUID)
	logger.Info(label, logger.Fields{
		"requestUuid": requestUuid,
	})
	return requestUuid
}