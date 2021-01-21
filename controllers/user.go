package controllers

import (
	"errors"
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
	userId, _, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	user := &sql.TbUser{
		UserId:      userId,
		Password:    s.Password,
		Phone:       s.Phone,
		StudentCard: s.StudentCard,
		PositionId:  s.PositionId,
		ClassId:     s.ClassId,
		Gender:      s.Gender,
		Residence:   s.Residence,
		Age:         s.Age,
		Email:       s.Email,
	}
	if err := logic.SetUp(user); err != nil {
		zap.L().Error("logic.SetUp(user) failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeNeedLogin)
		} else if errors.Is(err, logic.ErrorInvalidFormatOfPhone) || errors.Is(err, logic.ErrorInvalidFormatOfCard) {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		}
	} else {
		ResponseSuccess(c, nil)
	}
}

// UserDetailHandler 获取用户个人信息
func UserDetailHandler(c *gin.Context) {
	// 获取用户Id
	userId, _, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 业务逻辑
	data, err := logic.GetUserDetail(userId)
	if err != nil {
		zap.L().Error("logic.GetUserDetail(userId) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetAllPersonHandler 获取班级所有成员
func GetAllPersonHandler(c *gin.Context) {
	// 获取用户的权限
	userid, _, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	data, err := logic.GetAllPerson(userid)
	if err != nil {
		zap.L().Error("logic.GetAllPerson(userid) failed", zap.Error(err))
		if errors.Is(err, logic.ErrorNotInClass) {
			ResponseError(c, CodeNotInClass)
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseSuccess(c, data)
}

// GetUserNameHandler 获取用户名称
func GetUserNameHandler(c *gin.Context) {
	studentCard := c.Param("card")
	userid, positionId, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	name, err := logic.GetUserName(studentCard, userid, positionId)
	if err != nil {
		zap.L().Error("logic.GetUserName(studentCard, userid, positionId) failed", zap.Error(err))
		if errors.Is(err, logic.ErrorNotInClass) {
			ResponseError(c, CodeNotInClass)
		} else if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseErrorWithMsg(c, CodeUserNotExist, "该用户不在班级中或该用户不存在")
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseSuccess(c, name)
}

// SearchPersonHandler 搜索某个用户
func SearchPersonHandler(c *gin.Context) {
	// 参数解析
	s := new(models.ParamSearchUser)
	if err := c.ShouldBindJSON(s); err != nil {
		zap.L().Error("search person with invalid param", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	// 业务逻辑
	data, err := logic.SearchPerson(s)
	if err != nil {
		zap.L().Error("logic.SearchPerson(s) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// AddPersonHandler 教师添加学生
func AddPersonHandler(c *gin.Context) {
	// 参数解析
	s := new(models.ParamSearchUser)
	if err := c.ShouldBindJSON(s); err != nil {
		zap.L().Error("add person with invalid param", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	//	获取用户 id 和用户权限
	userId, positionId, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 只有班主任可以添加学生
	if positionId != 2 {
		ResponseError(c, CodeNotEnoughPermission)
		return
	}
	// 业务逻辑
	err = logic.AddPerson(s, userId)
	if err != nil {
		zap.L().Error("logic.AddPerson(s, userId) failed", zap.Error(err))
		if errors.Is(err, logic.ErrorInClass) {
			ResponseError(c, CodeInClass)
		} else if errors.Is(err, logic.ErrorNotInClass) {
			ResponseError(c, CodeNotInClass)
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseSuccess(c, nil)
}

// RemovePersonHandler 教师删除学生
func RemovePersonHandler(c *gin.Context) {
	// 参数解析
	s := new(models.ParamSearchUser)
	if err := c.ShouldBindJSON(s); err != nil {
		zap.L().Error("remove person with invalid param", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	//	获取用户 id 和用户权限
	userId, positionId, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 只有班主任可以删除学生
	if positionId != 2 {
		ResponseError(c, CodeNotEnoughPermission)
		return
	}
	err = logic.RemovePerson(s, userId)
	if err != nil {
		zap.L().Error("logic.RemovePerson(s, userId) failed", zap.Error(err))
		// 判断错误类型
		if errors.Is(err, logic.ErrorNotSameClass) {
			ResponseError(c, CodeNotSameClass)
		} else if errors.Is(err, logic.ErrorNotInClass) {
			ResponseError(c, CodeNotInClass)
		} else {
			ResponseError(c, CodeServerBusy)
		}
	} else {
		ResponseSuccess(c, nil)
	}
}

// UpdatePositionHandler 教师修改学生权限
func UpdatePositionHandler(c *gin.Context) {
	p := new(models.ParamUpdatePosition)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("update position with invalid param", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	//	获取用户 id 和用户权限
	userId, positionId, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 只有班主任可以删除学生
	if positionId != 2 {
		ResponseError(c, CodeNotEnoughPermission)
		return
	}
	err = logic.UpdatePosition(p, userId)
	if err != nil {
		zap.L().Error("logic.UpdatePosition(p, userId) failed", zap.Error(err))
		// 错误类型分析
		if errors.Is(err, logic.ErrorNotInClass) || errors.Is(err, logic.ErrorNotSameClass) || errors.Is(err, logic.ErrorPositionExist) {
			ResponseErrorWithMsg(c, CodeNotExpect, err.Error())
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseSuccess(c, nil)
}
