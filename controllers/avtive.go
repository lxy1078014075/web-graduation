package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
	"web-graduation/logic"
	"web-graduation/models"
)

// ActiveHandler 获取活动列表
func ActiveHandler(c *gin.Context) {
	userId, _, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if actives, err := logic.Active(userId); err != nil {
		zap.L().Error("logic.Active(userId) failed", zap.Error(err))
		if errors.Is(err, logic.ErrorNotInClass) {
			ResponseError(c, CodeNotInClass)
			return
		}
		ResponseError(c, CodeServerBusy)
	} else {
		ResponseSuccess(c, actives)
	}
}

// ActiveDetailHandler 获取活动详情
func ActiveDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetActiveDetail(id)
	if err != nil {
		zap.L().Error("logic.GetActiveDetail(id) failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, data)
}

// AddActiveHandler 新建活动
func AddActiveHandler(c *gin.Context) {
	active := new(models.ParamAddActive)
	if err := c.ShouldBindJSON(active); err != nil {
		zap.L().Error("add active with invalid param", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	// 获取用户ID
	userId, _, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 逻辑处理
	if err := logic.AddActive(active, userId); err != nil {
		zap.L().Error("logic.AddActive(active, userId) failed", zap.Error(err))
		if errors.Is(err, logic.ErrorNeedOptions) || errors.Is(err, logic.ErrorNotInClass) {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseError(c, CodeServerBusy)
	} else {
		ResponseSuccess(c, nil)
	}
}

// ModifyActiveHandler 修改活动
func ModifyActiveHandler(c *gin.Context)  {

}

// RemoveActiveHandler 删除活动
func RemoveActiveHandler(c *gin.Context)  {
	idStr:=c.Param("id")
	activeId,err:=strconv.ParseInt(idStr,10,64)
	if err!=nil{
		ResponseError(c,CodeInvalidParam)
		return
	}
	userId,positionId,err:=getCurrentUser(c)
	if err!=nil{
		ResponseError(c,CodeNeedLogin)
		return
	}
	err=logic.RemoveActive(activeId,userId,positionId)
	if err!=nil{
		zap.L().Error("logic.RemoveActive(activeId,userId) failed",zap.Error(err))
		// 错误类型处理
	}
	ResponseSuccess(c,nil)
}