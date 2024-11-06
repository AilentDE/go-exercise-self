package router

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func syncHandler(ctx *gin.Context) {
	time.Sleep(5 * time.Second)
	log.Println("syncHandler done", ctx.Request.URL.Path)

	ctx.JSON(200, gin.H{"message": "syncHandler"})
}

func asyncHandler(ctx *gin.Context) {
	cCp := ctx.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("asyncHandler done", cCp.Request.URL.Path)
	}()

	ctx.JSON(200, gin.H{"message": "asyncHandler"})
}