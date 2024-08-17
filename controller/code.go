package controller

// 定义程序可能出现的错误码

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassWord
	CodeServerBusy
	CodeNeedKLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassWord: "用户名或者密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeNeedKLogin:      "需要登录",
	CodeInvalidToken:    "无效的Token",
}

func (code ResCode) Msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
