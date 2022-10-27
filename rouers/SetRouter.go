package rouers

import (
	config "JoinUs/configs"
	"JoinUs/controller"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	pprof.Register(r)
	r.Use(config.LoggerToFile())

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
