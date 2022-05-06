package controllers

import "github.com/gin-gonic/gin"

func SourceControllers(engine *gin.Engine) {
	engine.GET("/", getRoot)
	engine.GET("/ping", getPing)
}
