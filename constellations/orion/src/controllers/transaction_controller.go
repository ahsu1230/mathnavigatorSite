package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetTransactionsByAccountId(c *gin.Context) {
	utils.LogControllerMethod(c, "transactionController.GetTransactionsByAccountId")
	accountId, err := utils.ParseParamId(c, "accountId")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("accountId")))
		c.Abort()
		return
	}

	transactionList, err := repos.TransactionRepo.SelectByAccountId(accountId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &transactionList)
}

func GetTransactionById(c *gin.Context) {
	utils.LogControllerMethod(c, "transactionController.GetTransactionById")
	//Incoming params
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	transaction, err := repos.TransactionRepo.SelectById(id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &transaction)
}

func CreateTransaction(c *gin.Context) {
	utils.LogControllerMethod(c, "transactionController.CreateTransaction")
	//JSON
	var transactionJson domains.Transaction
	if err := c.ShouldBindJSON(&transactionJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := transactionJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.TransactionRepo.Insert(transactionJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateTransaction(c *gin.Context) {
	utils.LogControllerMethod(c, "transactionController.UpdateTransaction")
	// Incoming JSON & Parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	var transactionJson domains.Transaction
	if err := c.ShouldBindJSON(&transactionJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := transactionJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	if err := repos.TransactionRepo.Update(id, transactionJson); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteTransaction(c *gin.Context) {
	utils.LogControllerMethod(c, "transactionController.DeleteTransaction")
	// Incoming Parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	if err := repos.TransactionRepo.Delete(id); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func GetAllPaymentTypes(c *gin.Context) {
	utils.LogControllerMethod(c, "transactionController.GetAllPaymentTypes")
	c.JSON(http.StatusOK, domains.ALL_TRANSACTION_TYPES)
}
