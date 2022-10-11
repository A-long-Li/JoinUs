package controller

import (
	"JoinUs/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertUser(ctx *gin.Context) {
	user := models.User{
		Name:       ctx.PostForm("name"),
		College:    ctx.PostForm("initials"),
		Major:      ctx.PostForm("top_domains"),
		ID:         ctx.PostForm("userNumber"),
		Qq:         ctx.PostForm("QQ"),
		Tel:        " ",
		Intention1: ctx.PostForm("purpose1"),
		Intention2: ctx.PostForm("purpose2"),
	}
	//查询是否已经提交过了
	err := models.FindUserById(&user)
	//之前已经提交过了,提醒覆盖
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "已覆盖您之前提交的信息",
		})
		err = models.UpdaterUser(&user)
	} else {
		if err = models.CreateUser(&user); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"msg":  err.Error(),
				"code": 502,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "恭喜你 报名成功",
			})
		}
	}

}
