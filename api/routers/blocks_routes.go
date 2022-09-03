package routers

import "github.com/gin-gonic/gin"

func SetBlocksRoutes(router *gin.RouterGroup) {
	router.GET("/:id", func(c *gin.Context) {})
	router.GET("", func(c *gin.Context) {})
}
