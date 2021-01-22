package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
	"web-graduation/dao/mysql"
	"web-graduation/logic"
	"web-graduation/models"
)

// VoteHandler 用户投票
func VoteHandler(c *gin.Context) {
	// 参数解析
	p := new(models.ParamVote)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("vote with invalid param", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	// 获取用户id
	userId, _, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	err = logic.Vote(p, userId)
	if err != nil {
		zap.L().Error("logic.Vote(p,userId) failed", zap.Error(err))
		// 错误类型的处理
		if errors.Is(err, mysql.ErrorNotExist) {
			ResponseErrorWithMsg(c, CodeNotExpect, err.Error())
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseSuccess(c, nil)
}

// VoteDetailHandler 获取当前的投票情况
func VoteDetailHandler(c *gin.Context) {
	// param参数的获取
	idStr := c.Param("id")
	activeId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.VoteDetail(activeId)
	if err != nil {
		zap.L().Error("logic.VoteDetail(activeId) failed", zap.Error(err))
		// 错误类型处理
		if errors.Is(err, logic.ErrorActiveHasProblem) {
			ResponseErrorWithMsg(c, CodeNotExpect, err.Error())
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseSuccess(c, data)
}

// VoteResultHandler 获取投票结果
func VoteResultHandler(c *gin.Context) {
	idStr := c.Param("id")
	activeId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.VoteResult(activeId)
	if err != nil {
		zap.L().Error("logic.VoteResult(activeId) failed", zap.Error(err))
		// 错误类型处理

		return
	}
	ResponseSuccess(c, data)
}
