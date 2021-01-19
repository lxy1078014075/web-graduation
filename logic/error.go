package logic

import "errors"

var (
	ErrorInvalidFormatOfPhone = errors.New("电话号码必须为11位的数字")
	ErrorInvalidFormatOfCard  = errors.New("学号必须为10位的数字")
	ErrorNeedOptions =errors.New("该活动需要选项")
)
