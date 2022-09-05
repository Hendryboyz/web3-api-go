package routers

import (
	"github.com/Hendryboyz/web3-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func SetBlocksRoutes(router *gin.RouterGroup) {
	controller := new(controllers.BlocksController)
	router.GET("/:id", controller.GetBlockById)
	router.GET("", controller.GetBlocks)
}
