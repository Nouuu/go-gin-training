package controllers

import "github.com/gin-gonic/gin"

func getPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
