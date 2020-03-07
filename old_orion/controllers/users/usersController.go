package users

//import (
//	"net/http"
//	"strconv"
//
//	"github.com/gin-gonic/gin"
//)
//
//func GetUsers(c *gin.Context) {
//	// Query Repo
//	userList := GetAllUsers()
//
//	// JSON Response
//	c.JSON(http.StatusOK, userList)
//	return
//}
//
//func GetUser(c *gin.Context) {
//	id := parseParamId(c, "id")
//
//	// Query Repo
//	user, err := GetUserById(id)
//	if err != nil {
//		panic(err)
//	} else {
//		c.JSON(http.StatusOK, user)
//	}
//	return
//}
//
//func CreateUser(c *gin.Context) {
//	// Incoming JSON
//	var userJson User
//	c.BindJSON(&userJson)
//
//	if err := CheckValidUser(userJson); err != nil {
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	// Query Repo (INSERT & SELECT)
//	err := InsertUser(userJson)
//	if err != nil {
//		panic(err)
//	} else {
//		c.Status(http.StatusNoContent)
//	}
//	return
//}
//
//func UpdateUser(c *gin.Context) {
//	// Incoming JSON & Parameters
//	id := parseParamId(c, "id")
//	var userJson User
//	c.BindJSON(&userJson)
//
//	if err := CheckValidUser(userJson); err != nil {
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	// Query Repo (UPDATE & SELECT)
//	err := UpdateUserById(id, userJson)
//	if err != nil {
//		panic(err)
//	} else {
//		c.Status(http.StatusNoContent)
//	}
//	return
//}
//
//func DeleteUser(c *gin.Context) {
//	id := parseParamId(c, "id")
//
//	// Query Repo (DELETE)
//	err := DeleteUserById(id)
//	if err != nil {
//		panic(err)
//	} else {
//		c.Status(http.StatusNoContent)
//	}
//	return
//}
//
//func parseParamId(c *gin.Context, key string) uint {
//	id, err := strconv.ParseUint(c.Param(key), 10, 32)
//	if err != nil {
//		panic(err)
//	}
//	return uint(id)
//}
//
//func CheckValidUser(user User) error {
//	/*
//	// Retrieves the inputted values
//	id := user.Id
//	firstName := user.FirstName
//	lastName := user.LastName
//	middleName := user.MiddleName
//	email := user.Email
//	phone := user.Phone
//	guardianId := user.GuardianId
//
//	// First name validation
//	// Last name validation
//	// Middle name validation
//	// Email validation
//	// Phone validation
//	// Guardian Id validation
//	if guardianId == id {
//		return errors.New("invalid guardian id")
//	}
//	*/
//	return nil
//}
//