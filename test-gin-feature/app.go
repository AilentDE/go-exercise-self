package main

import (
	"test-gin-feature/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	router.TestAsyncRoutes(server)
	router.TestBindUri(server)
	router.TestOnlyQuery(server)

	server.Run(":8080")
}