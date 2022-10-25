package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Hendryboyz/web3-api-go/configs"
	"github.com/Hendryboyz/web3-api-go/db"
	"github.com/Hendryboyz/web3-api-go/routers"
)

func main() {
	server := gin.New()

	configs := configs.ConfigServer("local")
	db.Connect()

	routers.InitRoutes(server)

	port := configs.GetInt32("server.port")

	server.Run(fmt.Sprintf("localhost:%d", port))
}
