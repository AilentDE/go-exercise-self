package router

import "github.com/gin-gonic/gin"

func TestAsyncRoutes(server *gin.Engine) {
	server.GET("/sync", syncHandler)
	server.GET("/async", asyncHandler)
}

func TestBindUri(server *gin.Engine) {
	server.GET("/user/:name/:id", bindUriHandler)
	server.GET("/user2/:name/:id", bindUriWithParamHandler)
}

func TestOnlyQuery(server *gin.Engine) {
	server.Any("/query", onlyQueryHandler)
}