package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义响应; 前后端分离时
/*
{
	"code": 10001. // 程序中的错误码
	“msg": xxx, //提示信息
	"date": {} // 存放数据
}
*/

const CtxUserIDKey = "userID"

type ResponseDate struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"` // omitempty 表示为空的时候不展示字段
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseDate{
		Code: code,
		Msg:  msg,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	re := &ResponseDate{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}

	c.JSON(http.StatusOK, re)
}

func ResponseError(c *gin.Context, code ResCode) {
	re := &ResponseDate{
		Code: code,
		Msg:  code.Msg(), // 返回错误的信息
		Data: nil,
	}

	c.JSON(http.StatusOK, re)
}
