package controllers

import (
	"gin-form/simple_music_api/data"
	"gin-form/simple_music_api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
		return
	}

	data.Albums = append(data.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
