package controller

import "C"
import (
	"encoding/json"
	"fmt"
	"ginDemo/cache"
	"ginDemo/constant"
	"ginDemo/model"
	"ginDemo/service"
	"ginDemo/util"
	"github.com/gin-gonic/gin"
)

func QueryUser(c *gin.Context) {
	result := util.Gin{C: c}
	userId := c.Param("userId")

	var user *model.UserInfo
	//从缓存获取
	cache := cache.Prefix{Name: constant.CACHE_USER}
	if cache.Exists(userId) {
		value, err := cache.Get(userId)
		if err != nil {
			result.Error(500, fmt.Sprintf(" err: %v", err))
		} else {
			json.Unmarshal(value, &user)
			result.Success(user)
		}
		return
	}

	service := service.UserInfo{UserId: userId}
	user, err := service.QueryUser()
	if err != nil {
		result.Error(500, fmt.Sprintf(" err: %v", err))
		return
	}
	//存储缓存
	cache.Set(user.UserId, result, 3600)
	result.Success(user)
}
