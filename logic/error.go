package logic

import "errors"

var (
	ErrorInvalidFormatOfPhone = errors.New("电话号码必须为11位的数字")
	ErrorInvalidFormatOfCard  = errors.New("学号必须为10位的数字")
	ErrorNeedOptions          = errors.New("该活动需要选项")
	ErrorNotInClass           = errors.New("没有分配班级，无法进行该操作")
	ErrorInClass              = errors.New("当前用户已经拥有班级，无法进行添加")
	ErrorNotSameClass         = errors.New("不在同一班级")
	ErrorPositionExist        = errors.New("该职位已经存在")
	ErrorActiveNotExist       = errors.New("该活动已不存在")
	ErrorNotSameCreator       = errors.New("您不是活动创始人，无法进行此操作")
)
