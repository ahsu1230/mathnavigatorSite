package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

// func GetAllUsers(c *gin.Context) {
// 	// Incoming optional parameter
// 	search := c.Query("search")
// 	pageSize := ParseParamInt(c.Query("pageSize"), 100)
// 	offset := ParseParamInt(c.Query("offset"), 0)

// 	userList, err := repos.UserRepo.SelectAll(search, pageSize, offset)
// 	if err != nil {
// 		c.Error(err)
// 		c.String(http.StatusInternalServerError, err.Error())
// 	} else {
// 		c.JSON(http.StatusOK, userList)
// 	}
// }

// type UserSearchBody struct {
// 	FirstName string
// 	MiddleName string
// 	LastName string
// 	Email string
// 	Phone string
// }

// func SearchUser(c *gin.Context) {
// 	// Incoming parameters
// 	var body UserSearchBody
// 	c.BindJSON(&body)

// 	FirstName := body.FirstName
// 	MiddleName := body.MiddleName
// 	LastName := body.LastName
// 	Email := body.Email
// 	Phone := body.Phone

// 	user, err := repos.UserRepo.SelectBy()
// 	if err != nil {
// 		c.Error(err)
// 		c.String(http.StatusNotFound, err.Error())
// 	} else {
// 		c.JSON(http.StatusOK, &user)
// 	}
// }

func GetUserById(c *gin.Context) {
	// Incoming parameters
	id := ParseParamId(c)

	user, err := repos.UserRepo.SelectById(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &user)
	}
}

func GetUsersByFamilyId(c *gin.Context) {
	// Incoming parameters
	familyId := ParseParamUint(c.Param("familyId"))

	user, err := repos.UserRepo.SelectByFamilyId(familyId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &user)
	}
}

func CreateUser(c *gin.Context) {
	// Incoming JSON
	var userJson domains.User
	c.BindJSON(&userJson)

	if err := userJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.UserRepo.Insert(userJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func UpdateUser(c *gin.Context) {
	// Incoming JSON & Parameters
	id := ParseParamId(c)
	var userJson domains.User
	c.BindJSON(&userJson)

	if err := userJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.UserRepo.Update(id, userJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteUser(c *gin.Context) {
	// Incoming Parameters
	id := ParseParamId(c)

	err := repos.UserRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
