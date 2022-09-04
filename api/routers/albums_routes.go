package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Hendryboyz/web3-api-go/controllers"
)

func SetAlbumsRoutes(router *gin.RouterGroup) {
	albumController := new(controllers.AlbumsController)
	router.GET("", albumController.GetAlbums)
	router.GET(":id", albumController.GetAlbumById)
	router.POST("", albumController.CreateAlbum)
}
