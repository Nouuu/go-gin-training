package controllers

import (
	"gin-form/simple_music_api/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range data.Albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}
