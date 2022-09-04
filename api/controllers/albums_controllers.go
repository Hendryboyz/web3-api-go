package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type AlbumsController struct{}

func (c AlbumsController) GetAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func (c AlbumsController) CreateAlbum(context *gin.Context) {
	var newAlbum album

	// invoke BindJSON() to bind(serialize) the request body to newAlbum variable
	if error := context.BindJSON(&newAlbum); error != nil {
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func (c AlbumsController) GetAlbumById(context *gin.Context) {
	id := context.Param("id")
	for _, eachAlbum := range albums {
		if eachAlbum.ID == id {
			context.IndentedJSON(http.StatusOK, eachAlbum)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
