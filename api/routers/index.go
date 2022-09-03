package routers

import "github.com/gin-gonic/gin"

// Setting routes
func InitRoutes(engine *gin.Engine) {
	SetHealthcheckRoute(engine)

	blocksGroup := engine.Group("/blocks")
	SetBlocksRoutes(blocksGroup)
	transactionGroup := engine.Group("/transaction")
	SetTransactionsRoutes(transactionGroup)

	version1 := engine.Group("/v1")
	albumGroup := version1.Group("/albums")
	SetAlbumsRoutes(albumGroup)
}
