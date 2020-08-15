package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

type AccountSearchBody struct {
	PrimaryEmail string
}

func GetAccountById(c *gin.Context) {
	// Incoming parameters
	id := ParseParamId(c)

	account, err := repos.AccountRepo.SelectById(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &account)
	}
}

func SearchAccount(c *gin.Context) {
	// Incoming parameters
	var body AccountSearchBody
	c.BindJSON(&body)
	primaryEmail := body.PrimaryEmail

	account, err := repos.AccountRepo.SelectByPrimaryEmail(primaryEmail)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &account)
	}
}

func GetNegativeBalanceAccounts(c *gin.Context) {
	accountSum, err := repos.AccountRepo.SelectAllNegativeBalances()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, accountSum)
	}
}

func CreateAccount(c *gin.Context) {
	// Incoming JSON
	var accountJson domains.Account
	c.BindJSON(&accountJson)

	if err := accountJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.AccountRepo.Insert(accountJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func UpdateAccount(c *gin.Context) {
	// Incoming JSON & Parameters
	id := ParseParamId(c)
	var accountJson domains.Account
	c.BindJSON(&accountJson)

	if err := accountJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.AccountRepo.Update(id, accountJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteAccount(c *gin.Context) {
	// Incoming Parameters
	id := ParseParamId(c)

	err := repos.AccountRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
