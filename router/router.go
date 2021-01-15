package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"web-graduation/controllers"
	"web-graduation/logger"
)

// Init 初始化路由
func Init() *gin.Engine {

	if err := controllers.InitTrans("zh"); err != nil {
		zap.L().Error("init validator trans failed", zap.Error(err))
	}

	r := gin.New()
	// 设置路由的两个默认中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1:=r.Group("/api/v1")

	// 注册账号
	v1.POST("/signup",controllers.SignUpHandler)
	// 登陆
	v1.POST("/login",controllers.LoginHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})
	return r
}
