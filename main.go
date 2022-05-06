package main

import (
	"go-gin-training/reverse_proxy_load_balancer"
)

func main() {
	reverse_proxy_load_balancer.StartServer(8080)
}
