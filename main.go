package main

import (
	"fmt"
	"ginDemo/cache"
	"ginDemo/model"
	"ginDemo/routers"
	"ginDemo/setting"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Init()
	model.Init()
	cache.Init()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	r := routers.Init()

	r.Run(fmt.Sprintf(":%d", setting.ServerSetting.HttpPort))
}
