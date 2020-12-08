package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

func retrieveRequestUuid(c *gin.Context) (uuid.UUID, bool) {
	v, ok := c.Get(domains.REQUEST_UUID)
	if ok {
		requestUuid := v.(uuid.UUID)
		return requestUuid, true
	}
	return uuid.UUID{}, false
}

func LogControllerMethod(c *gin.Context, label string) {
	requestUuid, _ := retrieveRequestUuid(c)
	logger.Debug(label, logger.Fields{
		"requestUuid": requestUuid,
	})
}

func RetrieveContext(c *gin.Context) context.Context {
	ctx := c.Request.Context()
	requestUuid, ok := retrieveRequestUuid(c)
	if !ok {
		return ctx
	}
	return context.WithValue(ctx, domains.REQUEST_UUID, requestUuid)
}
