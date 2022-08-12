package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Success(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
	return
}

func (g *Gin) Error(errCode int, errMsg string) {
	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  errMsg,
	})
	return
}
