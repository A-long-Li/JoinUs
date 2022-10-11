package controller

import "github.com/gin-gonic/gin"

func GetIndex(ctx *gin.Context) {
	ctx.HTML(200, "index-1.html", nil)
}
