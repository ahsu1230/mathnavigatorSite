package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetTransactionsByAccountId(c *gin.Context) {
	accountId, _ := utils.ParseParamIdString(c, "accountId")

	transactionList, err := repos.TransactionRepo.SelectByAccountId(accountId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, &transactionList)
}

func GetTransactionById(c *gin.Context) {
	//Incoming params
	id, _ := utils.ParseParamIdString(c, "id")

	transaction, err := repos.TransactionRepo.SelectById(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, &transaction)
}

func CreateTransaction(c *gin.Context) {
	//JSON
	var transactionJson domains.Transaction
	c.BindJSON(&transactionJson)

	if err := transactionJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.TransactionRepo.Insert(transactionJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func UpdateTransaction(c *gin.Context) {
	// Incoming JSON & Parameters
	id, _ := utils.ParseParamIdString(c, "id")
	var transactionJson domains.Transaction
	c.BindJSON(&transactionJson)

	if err := transactionJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.TransactionRepo.Update(id, transactionJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func DeleteTransaction(c *gin.Context) {
	// Incoming Parameters
	id, _ := utils.ParseParamIdString(c, "id")

	err := repos.TransactionRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func GetAllPaymentTypes(c *gin.Context) {
	c.JSON(http.StatusOK, domains.ALL_TRANSACTION_TYPES)
}
