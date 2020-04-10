package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
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

func ParseParamNullUint(nullUint string) domains.NullUint {
	id, err := strconv.ParseUint(nullUint, 10, 32)
	if err != nil {
		panic(err)
	}
	return domains.NewNullUint(uint(id))
}
