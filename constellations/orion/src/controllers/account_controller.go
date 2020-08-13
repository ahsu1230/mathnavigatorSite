package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
)

type AccountSearchBody struct {
	PrimaryEmail string
}

func GetAccountById(c *gin.Context) {
	utils.LogControllerMethod(c, "accountController.GetAccountById")

	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		err = appErrors.WrapParse(err, c.Param("id"))
		c.Error(err)
		c.Abort()
		return
	}

	account, err := repos.AccountRepo.SelectById(id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &account)
}

func SearchAccount(c *gin.Context) {
	utils.LogControllerMethod(c, "accountController.SearchAccount")

	var body AccountSearchBody
	if err := c.ShouldBindJSON(&body); err != nil {
		err = appErrors.WrapBindJSON(err, c.Request)
		c.Error(err)
		c.Abort()
		return
	}
	primaryEmail := body.PrimaryEmail

	account, err := repos.AccountRepo.SelectByPrimaryEmail(primaryEmail)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &account)
}

func CreateAccount(c *gin.Context) {
	utils.LogControllerMethod(c, "accountController.CreateAccount")

	var accountJson domains.Account
	if err := c.BindJSON(&accountJson); err != nil {
		err = appErrors.WrapBindJSON(err, c.Request)
		c.Error(err)
		c.Abort()
		return
	}

	if err := accountJson.Validate(); err != nil {
		err = appErrors.WrapInvalidDomain(err, "Invalid Account")
		c.Error(err)
		c.Abort()
		return
	}

	if err := repos.AccountRepo.Insert(accountJson); err != nil {
		err = appErrors.WrapRepo(err)
		c.Error(err)
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateAccount(c *gin.Context) {
	utils.LogControllerMethod(c, "accountController.UpdateAccount")

	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		err = appErrors.WrapParse(err, c.Param("id"))
		c.Error(err)
		c.Abort()
		return
	}

	var accountJson domains.Account
	if err = c.BindJSON(&accountJson); err != nil {
		err = appErrors.WrapBindJSON(err, c.Request)
		c.Error(err)
		c.Abort()
		return
	}

	if err = accountJson.Validate(); err != nil {
		err = appErrors.WrapInvalidDomain(err, "Invalid Account")
		c.Error(err)
		c.Abort()
		return
	}

	if err = repos.AccountRepo.Update(id, accountJson); err != nil {
		err = appErrors.WrapRepo(err)
		c.Error(err)
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteAccount(c *gin.Context) {
	utils.LogControllerMethod(c, "accountController.DeleteAccount")

	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		err = appErrors.WrapParse(err, c.Param("id"))
		c.Error(err)
		c.Abort()
		return
	}

	if err := repos.AccountRepo.Delete(id); err != nil {
		err = appErrors.WrapRepo(err)
		c.Error(err)
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}
