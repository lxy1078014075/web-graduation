package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web-graduation/dao/mysql"
	"web-graduation/logic"
	"web-graduation/models"
)

func SignUpHandler(c *gin.Context)  {
	// 1. 获取参数、参数校验
	p:=new(models.ParamSignUp)
	if err:=c.ShouldBindJSON(&p);err!=nil{
		zap.L().Error("SignUp with invalid param",zap.Error(err))
		// 获取validator.ValidationErrors类型的errors
		errs,ok:=err.(validator.ValidationErrors)
		if !ok{
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c,CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParam,removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 注册业务处理
	if err:=logic.SignUp(p);err!=nil{
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

func LoginHandler(c *gin.Context)  {

}