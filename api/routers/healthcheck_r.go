package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func SetHealthcheckRoute(engine *gin.Engine) {
	engine.GET("/health", healthCheck)
}
