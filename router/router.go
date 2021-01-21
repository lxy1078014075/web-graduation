package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"web-graduation/controllers"
	"web-graduation/logger"
	"web-graduation/middlewares"
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
	// 登陆获取 token
	v1.POST("/login",controllers.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())	//应用jwt中间件
	{
		// 获取用户的个人信息 (教师部分善未完善)
		v1.GET("/user_info",controllers.UserDetailHandler)
		// 修改用户信息
		v1.POST("/setup",controllers.SetUpHandler)
		// 获取班级的所有成员
		v1.GET("/all_person",controllers.GetAllPersonHandler)
		// 获取活动列表
		v1.GET("/active",controllers.ActiveHandler)
		// 获取活动详情
		v1.GET("/active/:id",controllers.ActiveDetailHandler)
		// 新增活动
		v1.POST("/add_active",controllers.AddActiveHandler)
	}



	// 路由不存在时的返回
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})
	return r
}
