package controller

import "github.com/gin-gonic/gin"

func Register(ctx *gin.Context) {
	ctx.HTML(200, "login-1.html", nil)
}
