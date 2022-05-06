package static_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer(port int) {
	router := gin.Default()

	router.Static("/", "./static_server/static")
	router.NoRoute(func(c *gin.Context) {
		c.File("./static_server/static/index.html")
	})

	router.Run(fmt.Sprintf(":%d", port))
}
