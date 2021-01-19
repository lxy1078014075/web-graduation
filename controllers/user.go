package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web-graduation/dao/mysql"
	"web-graduation/logic"
	"web-graduation/models"
	"web-graduation/models/sql"
)

// SignUpHandler 注册控制器
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数、参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 注册业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("SignUp into database failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登陆控制器
func LoginHandler(c *gin.Context) {
	// 参数获取与校验
	login := new(models.ParamLogin)
	if err := c.ShouldBindJSON(login); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	// 业务逻辑
	if token, err := logic.Login(login); err != nil {
		zap.L().Error("logic.Login(login) failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeServerBusy)
	} else {
		ResponseSuccess(c, token)
	}
}

// SetUpHandler 设置用户个人信息
func SetUpHandler(c *gin.Context) {
	// 参数校验
	s := new(models.ParamSetUp)
	if err := c.ShouldBindJSON(s); err != nil {
		zap.L().Error("setup with invalid param", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	// 业务逻辑
	userId, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	fmt.Println(userId)
	user := &sql.TbUser{
		UserId:      userId,
		Password:    s.Password,
		Phone:       s.Phone,
		StudentCard: s.StudentCard,
	}
	if err := logic.SetUp(user); err != nil {
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeNeedLogin)
		} else if errors.Is(err,logic.ErrorInvalidFormatOfPhone) || errors.Is(err,logic.ErrorInvalidFormatOfCard) {
			ResponseErrorWithMsg(c,CodeInvalidParam,err.Error())
		}
	} else {
		ResponseSuccess(c, nil)
	}
}