package routers

import (
	"github.com/Hendryboyz/web3-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func SetTransactionsRoutes(router *gin.RouterGroup) {
	controller := new(controllers.TransactionsController)
	router.GET("/:txHash", controller.GetTransactionByHash)
}
