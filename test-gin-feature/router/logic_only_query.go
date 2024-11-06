package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
}

func onlyQueryHandler(ctx *gin.Context) {
	var person Person

	if ctx.ShouldBindQuery(&person) == nil {
		log.Println("Name", person.Name)
		log.Println("Address", person.Address)
	}

	ctx.JSON(200, gin.H{"message": "onlyQueryHandler"})
}