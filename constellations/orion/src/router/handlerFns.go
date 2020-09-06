package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
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
			"requestUuid":   requestUuid,
			"requestMethod": c.Request.Method,
			"requestURL":    c.Request.URL,
			"requestHost":   c.Request.Host,
		})
		c.Set(domains.REQUEST_UUID, requestUuid)
		c.Writer.Header().Set("X-Request-Id", requestUuid.String())

		c.Next()

		// After request finishes, prepare response
		if len(c.Errors) > 0 {
			respError := createAppErrorFromResponseErrors(c)
			logger.Error(respError.Message, respError.Error, logger.Fields{
				"requestUuid": requestUuid,
				"code":        respError.Code,
			})
			c.AbortWithStatusJSON(respError.Code, &respError)
			return
		}

		if !c.IsAborted() {
			logger.Info("Succesfully completed request!", logger.Fields{
				"requestUuid": requestUuid,
				"fullPath":    c.FullPath(),
				"status":      c.Writer.Status(),
			})
		}
	}
}

func createAppErrorFromResponseErrors(c *gin.Context) appErrors.ResponseError {
	// For now, assume only one Error collected
	wrappedErr := c.Errors[0].Err

	logger.Error("Handling error from request", wrappedErr, logger.Fields{
		"cause": errors.Cause(wrappedErr),
	})

	var message string
	var code int

	if errors.Is(wrappedErr, appErrors.ERR_INVALID_DOMAIN) {
		message = wrappedErr.Error()
		code = http.StatusBadRequest

	} else if errors.Is(wrappedErr, appErrors.ERR_JSON_NULL_BODY) {
		message = "Must provide a JSON body for this request"
		code = http.StatusBadRequest

	} else if errors.Is(wrappedErr, appErrors.ERR_JSON_BIND_BODY) {
		message = "Could not bind JSON body for this request. Please provide a valid JSON or domain."
		code = http.StatusBadRequest

	} else if errors.Is(wrappedErr, appErrors.ERR_REPO_EXEC_MISMATCH) {
		message = "Database update was not saved due to result mismatch."
		code = http.StatusInternalServerError

	} else if errors.Is(wrappedErr, appErrors.ERR_MYSQL_DUPLICATE_ENTRY) {
		message = "Database duplicate entry conflict. Please change some fields."
		code = http.StatusBadRequest

	} else if errors.Is(wrappedErr, appErrors.ERR_SQL_NO_ROWS) {
		message = "No results from query. Please change your search terms."
		code = http.StatusNotFound

	} else {
		message = "Unknown error"
		code = http.StatusNotImplemented // (501)
	}

	return appErrors.ResponseError{
		Code:    code,
		Message: message,
		Error:   wrappedErr,
	}
}
