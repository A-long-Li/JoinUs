package rouers

import (
	"JoinUs/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	{
		r.Static("/static", "./static")
		r.LoadHTMLGlob("static/template/*")
	}

	{
		r.GET("/", controller.GetIndex)
		r.GET("/index-1.html", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/")
		})
	}
	{
		r.GET("/register", controller.Register)
		r.GET("/login-1.html", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/register")
		})
	}
	{
		r.POST("/register/add", controller.InsertUser)
	}
	return r
}
