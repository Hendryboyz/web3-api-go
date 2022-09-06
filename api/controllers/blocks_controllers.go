package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlocksController struct{}

func (c *BlocksController) GetBlockById(context *gin.Context) {
	blockId := context.Param("id")
	context.IndentedJSON(http.StatusOK, gin.H{"id": blockId})
}

func (c *BlocksController) GetBlocks(context *gin.Context) {
	limit := strToInt64(context.Query("limit"))
	context.IndentedJSON(http.StatusOK, gin.H{"limit": limit})
}

func strToInt64(s string) int64 {
	number, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	} else {
		return number
	}
}
