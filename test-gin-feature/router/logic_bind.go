package router

import "github.com/gin-gonic/gin"

type User struct {
	Name string `uri:"name" binding:"required"`
	Id   string `uri:"id" binding:"required"`
}

func bindUriHandler(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"name": user.Name, "id": user.Id})
}

func bindUriWithParamHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{"name": name, "id": id})
}
