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

func ParseParamIds(c *gin.Context, key string) []uint {
	ids := make([]uint, 0)

	for _, id := range c.PostFormArray(key) {
		id, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			panic(err)
		}
		ids = append(ids, uint(id))
	}
	return ids
}
