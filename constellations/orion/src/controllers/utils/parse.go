package utils

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func ParseParamIdString(c *gin.Context, idString string) (uint, error) {
	id, err := strconv.ParseUint(c.Param(idString), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func ParseParamUint(numStr string) (uint, error) {
	id, err := strconv.ParseUint(numStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func ParseParamInt(str string, preset int) (int, error) {
	if len(str) == 0 {
		return preset, nil
	}

	integer, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(integer), nil
}

func ParseParamPublishedOnly(c *gin.Context) bool {
	return c.Query("published") == "true"
}
