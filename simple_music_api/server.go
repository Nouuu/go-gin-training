package simple_music_api

import (
	"fmt"
	"gin-form/simple_music_api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer(port int) {
	router := gin.Default()

	controllers.SourceControllers(router)

	router.Run(fmt.Sprintf(":%d", port))
}
