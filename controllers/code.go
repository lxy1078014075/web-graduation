package controllers

type Code int

const (
	CodeSuccess Code = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidHeader
	CodeInvalidToken

	CodeRedisNotExist
)

var codeMsgMap = map[Code]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "账号或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeNeedLogin:     "请登录",
	CodeInvalidHeader: "无效的请求头",
	CodeInvalidToken:  "无效的token",

	CodeRedisNotExist: "Redis不存在",
}

func (c Code) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
