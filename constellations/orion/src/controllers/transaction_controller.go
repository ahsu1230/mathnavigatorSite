package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	transactionList, err := repos.TransactionRepo.SelectAll()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, transactionList)
	}
}

func GetTransactionById(c *gin.Context) {
	//Incoming params
	id := ParseParamId(c)

	transaction, err := repos.TransactionRepo.SelectById(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &transaction)
	}
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
	} else {
		c.Status(http.StatusOK)
	}
}

func UpdateTransaction(c *gin.Context) {
	// Incoming JSON & Parameters
	id := ParseParamId(c)
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
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteTransaction(c *gin.Context) {
	// Incoming Parameters
	id := ParseParamId(c)

	err := repos.TransactionRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
