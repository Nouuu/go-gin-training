package reverse_proxy_load_balancer

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer(port int) {
	router := gin.Default()

	router.Any("/*any", Proxy)

	router.Run(fmt.Sprintf(":%d", port))
}
