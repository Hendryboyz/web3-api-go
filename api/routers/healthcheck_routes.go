package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetHealthcheckRoute(engine *gin.Engine) {
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
}
