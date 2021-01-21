package controllers

type Code int

const (
	CodeSuccess Code = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNotEnoughPermission
	CodeParamIsNil

	CodeNeedLogin
	CodeInvalidHeader
	CodeInvalidToken

	CodeRedisNotExist

	CodeNotInClass
	CodeInClass
	CodeNotSameClass

	CodeNotExpect	// 不应该出现的错误
)

var codeMsgMap = map[Code]string{
	CodeSuccess:             "success",
	CodeInvalidParam:        "请求参数错误",
	CodeUserExist:           "用户已存在",
	CodeUserNotExist:        "用户不存在",
	CodeInvalidPassword:     "账号或密码错误",
	CodeServerBusy:          "服务繁忙",
	CodeNotEnoughPermission: "用户权限不够",
	CodeParamIsNil:          "参数都为空",

	CodeNeedLogin:     "请登录",
	CodeInvalidHeader: "无效的请求头",
	CodeInvalidToken:  "无效的token",

	CodeRedisNotExist: "Redis不存在",
	CodeNotInClass:    "没有分配班级，无法进行该操作",
	CodeInClass:       "当前用户已经拥有班级，无法进行添加",
	CodeNotSameClass: "当前用户不在同一班级",
}

func (c Code) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
