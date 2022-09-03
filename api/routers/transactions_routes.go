package routers

import "github.com/gin-gonic/gin"

func SetTransactionsRoutes(router *gin.RouterGroup) {
	router.GET("/:txHash", func(c *gin.Context) {})
}
