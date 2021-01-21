package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"
const CtxUserPositionIdKey = "positionId"

var ErrorUserNotLogin = errors.New("用户未登陆")

func getCurrentUser(c *gin.Context) (userId int64, positionId int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	positionId = c.GetInt64(CtxUserPositionIdKey)
	return
}
