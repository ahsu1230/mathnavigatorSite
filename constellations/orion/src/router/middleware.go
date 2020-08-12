package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "github.com/pkg/errors"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("API endpoint unrecognized by handler", logger.Fields{})
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Unrecognized path",
		})
	}
}

func AppRequestHandler() gin.HandlerFunc {
    return func(c *gin.Context) {

		// Log with Request UUID
		requestUuid := uuid.New()
		logger.Info("Received Request", logger.Fields{
			"requestUuid": requestUuid,
			"requestMethod": c.Request.Method,
			"requestURL": c.Request.URL,
			"requestHost": c.Request.Host,
		})
		c.Set("requestUuid", requestUuid)
		c.Writer.Header().Set("X-Request-Id", requestUuid.String())

		c.Next()

		// After request finishes, prepare response
		if (len(c.Errors) > 0) {
			respError := createAppErrorFromResponseErrors(c)
			logger.Error(respError.Message, respError.Error, logger.Fields{
				"requestUuid": requestUuid,
				"code": respError.Code,
			})
			c.AbortWithStatusJSON(respError.Code, &respError)
			return
		}

		if (!c.IsAborted()) {
			logger.Info("Succesfully completed request!", logger.Fields{
				"requestUuid": requestUuid,
				"fullPath": c.FullPath(),
				"status": c.Writer.Status(),
			})
		}
	}
}

func createAppErrorFromResponseErrors(c *gin.Context) appErrors.ResponseError {
	// collectedErrors := c.Errors
	// get first one
	// if (err.Cause() == errors.constant), message = _, code = _
	return appErrors.ResponseError {}
}