package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Hendryboyz/web3-api-go/routers"
)

func main() {
	server := gin.New()

	// Setting routes
	routers.SetHealthcheckRoute(server)
	version1 := server.Group("/v1")
	routers.InitRoutes(version1)

	server.Run("localhost:8080")
}
