package routers

import (
	"ginDemo/controller"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/user/:userId", controller.QueryUser)
	return r
}
