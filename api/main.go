package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Hendryboyz/web3-api-go/configs"
	"github.com/Hendryboyz/web3-api-go/db"
	"github.com/Hendryboyz/web3-api-go/routers"
)

func main() {
	server := gin.New()

	configs.ConfigServer("local")
	db.Connect()

	routers.InitRoutes(server)

	server.Run("localhost:8080")
}
