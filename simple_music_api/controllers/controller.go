package controllers

import "github.com/gin-gonic/gin"

func SourceControllers(router *gin.Engine) {
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", addAlbum)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
