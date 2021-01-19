package logic

import (
	"fmt"
	"regexp"
	"web-graduation/dao/mysql"
	"web-graduation/models"
	"web-graduation/models/sql"
	"web-graduation/pkg/jwt"
	"web-graduation/pkg/snowflake"
)

// SignUp 注册的业务逻辑
func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if err = mysql.CheckUserExist(p.StudentCard); err != nil {
		return err
	}
	// 判断学号是否符合类型
	if p.StudentCard != "" && !VerifyCardFormat(p.StudentCard) {
		return ErrorInvalidFormatOfCard
	}
	// 生成ID
	userId := snowflake.GetID()
	user := &sql.TbUser{
		UserId:      userId,
		UserName:    p.Username,
		Password:    p.Password,
		StudentCard: p.StudentCard,
	}
	return mysql.InsertUser(user)
}

// Login 用户登陆逻辑
func Login(l *models.ParamLogin) (token string, err error) {
	user := &sql.TbUser{
		UserName: l.Username,
		Password: l.Password,
	}
	if err := mysql.LoginUser(user); err != nil {
		return "", err
	}
	fmt.Println(user.UserId, user.Identity)
	return jwt.GetToken(user.UserId, user.Identity)
}

// SetUp 设置用户信息逻辑
func SetUp(u *sql.TbUser) (err error) {
	// 判断电话是否符合类型
	if u.Phone != "" && !VerifyMobileFormat(u.Phone) {
		return ErrorInvalidFormatOfPhone
	}
	// 判断学号是否符合类型
	if u.StudentCard != "" && !VerifyCardFormat(u.StudentCard) {
		return ErrorInvalidFormatOfCard
	}
	return mysql.SetUpUser(u)
}

// VerifyMobileFormat 正则表达式判断是否是手机号
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// VerifyCardFormat 正则表达式判断是否是学号
func VerifyCardFormat(studentCard string) bool {
	regular := "^\\d{10}"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(studentCard)
}
