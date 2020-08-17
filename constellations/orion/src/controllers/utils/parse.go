package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseParamId(c *gin.Context, idString string) (uint, error) {
	id, err := strconv.ParseUint(c.Param(idString), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func ParseParamPublishedOnly(c *gin.Context) bool {
	return c.Query("published") == "true"
}
