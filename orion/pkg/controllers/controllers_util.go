package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseParamId(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(id)
}

func ParseParamPublishedOnly(c *gin.Context) bool {
	return c.Query("published") == "true"
}

func ParseParamUint(num string) uint {
	id, err := strconv.ParseUint(num, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(id)
}

func ParseParamInt(str string, preset int) int {
	if len(str) == 0 {
		return preset
	}

	integer, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(integer)
}
