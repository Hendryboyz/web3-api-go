package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionsController struct{}

func (c *TransactionsController) GetTransactionByHash(context *gin.Context) {
	hash := context.Param("txHash")
	context.IndentedJSON(http.StatusOK, gin.H{"txHash": hash})
}
