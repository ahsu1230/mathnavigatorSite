package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
)

type AccountSearchBody struct {
	PrimaryEmail string
}

func GetAccountById(c *gin.Context) {
	id, err := ParseParamIdString(c, "id")
	if err != nil {
		err = errors.Wrap(err, domains.ErrParse("id"))
		logger.Error("Error parsing id", err, logger.Fields{})
		c.Error(err)
		return
	}

	account, err := repos.AccountRepo.SelectById(id)
	if err != nil {
		err = errors.Wrap(err, domains.ErrRepo("account"))
		logger.Error("Error retrieving account", err, logger.Fields{ "id": id})
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, &account)
}

func SearchAccount(c *gin.Context) {
	var body AccountSearchBody
	if err := c.BindJSON(&body); err != nil {
		err = errors.Wrap(err, domains.ErrBindJson("AccountSearchBody"))
		logger.Error("Error binding incoming account search request", err, logger.Fields{})
		c.Error(err)
		return
	}
	primaryEmail := body.PrimaryEmail

	account, err := repos.AccountRepo.SelectByPrimaryEmail(primaryEmail)
	if err != nil {
		err = errors.Wrap(err, domains.ErrRepo("account"))
		logger.Error("Error retrieving account", err, logger.Field{
			"primaryEmail": primaryEmail,
		})
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, &account)
}

func CreateAccount(c *gin.Context) {
	var accountJson domains.Account
	if err := c.BindJSON(&body); err != nil {
		err = errors.Wrap(err, domains.ErrBindJson("Account"))
		logger.Error("Error deserializing create account request", err, logger.Fields{})
		c.Error(err)
		return
	}

	if err := accountJson.Validate(); err != nil {
		err = errors.Wrap(err, domains.ErrDomain("account"))
		logger.Error("Invalid account", err, logger.Fields{"account": accountJson})
		c.Error(err)
		return
	}
	
	if err := repos.AccountRepo.Insert(accountJson); err != nil {
		err = errors.Wrap(err, domains.ErrDomain("account"))
		logger.Error("Error inserting to account repo", err, logger.Fields{"account": accountJson})
		c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}

func UpdateAccount(c *gin.Context) {
	id, err := ParseParamIdString(c, "id")
	if err != nil {
		err = errors.Wrap(err, domains.ErrParse("id"))
		logger.Error("Error parsing id", err, logger.Fields{})
		c.Error(err)
		return
	}

	var accountJson domains.Account
	if err = c.BindJSON(&accountJson); err != nil {
		err = errors.Wrap(err, domains.ErrBind("Account"))
		logger.Error("Error binding account update request", err, logger.Fields{})
		c.Error(err)
		return
	}

	if err := accountJson.Validate(); err != nil {
		err = errors.Wrap(err, domains.ErrDomain("account"))
		logger.Error("Invalid account", err, logger.Fields{"account": accountJson})
		c.Error(err)
		return
	}

	err := repos.AccountRepo.Update(id, accountJson)
	if err != nil {
		err = errors.Wrap(err, domains.ErrRepo("account"))
		logger.Error("Error updating to account repo", err, logger.Fields{"account": accountJson})
		c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}

func DeleteAccount(c *gin.Context) {
	// Incoming Parameters
	id, err := ParseParamIdString(c, "id")
	if err != nil {
		err = errors.Wrap(err, domains.ErrParse("id"))
		logger.Error("Error parsing id", err, logger.Fields{})
		c.Error(err)
		return
	}

	if err := repos.AccountRepo.Delete(id); err != nil {
		err = errors.Wrap(err, domains.ErrRepo("account"))
		logger.Error("Error deleting from repo", err, logger.Fields{"id": id})
		c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}

func asdf(err error) domains.AppError {
	err = errors.Wrap(err, ...)
	logger.Error()
	return AppError {

	}
}