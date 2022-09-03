package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Hendryboyz/web3-api-go/routers"
)

func main() {
	server := gin.New()

	routers.InitRoutes(server)

	server.Run("localhost:8080")
}
