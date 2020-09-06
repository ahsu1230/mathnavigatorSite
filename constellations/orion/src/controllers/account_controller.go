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
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	account, err := repos.AccountRepo.SelectById(ctx, id)
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
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}
	primaryEmail := body.PrimaryEmail

	ctx := utils.RetrieveContext(c)
	account, err := repos.AccountRepo.SelectByPrimaryEmail(ctx, primaryEmail)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &account)
}

func GetNegativeBalanceAccounts(c *gin.Context) {
	ctx := utils.RetrieveContext(c)
	accountSum, err := repos.AccountRepo.SelectAllNegativeBalances(ctx)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, accountSum)
	}
}

func CreateAccountAndUser(c *gin.Context) {
	utils.LogControllerMethod(c, "accountController.CreateAccount")

	// Incoming JSON
	var accountUser domains.AccountUser
	if err := c.ShouldBindJSON(&accountUser); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}
	account := accountUser.Account
	user := accountUser.User

	if err := account.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}
	if err := user.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}
	ctx := utils.RetrieveContext(c)
	if err := repos.AccountRepo.InsertWithUser(ctx, account, user); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateAccount(c *gin.Context) {
	utils.LogControllerMethod(c, "accountController.UpdateAccount")

	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	var accountJson domains.Account
	if err = c.ShouldBindJSON(&accountJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err = accountJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	if err = repos.AccountRepo.Update(ctx, id, accountJson); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteAccount(c *gin.Context) {
	utils.LogControllerMethod(c, "accountController.DeleteAccount")

	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	if err := repos.AccountRepo.Delete(ctx, id); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
