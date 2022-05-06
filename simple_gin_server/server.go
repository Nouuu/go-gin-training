package simple_gin_server

import (
	"gin-form/simple_gin_server/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	engine := gin.Default()

	controllers.SourceControllers(engine)

	engine.Run(":9090")
}
