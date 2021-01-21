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
		StudentCard: l.StudentCard,
		Password:    l.Password,
	}
	if err := mysql.LoginUser(user); err != nil {
		return "", err
	}
	fmt.Println(user.UserId, user.PositionId)
	return jwt.GetToken(user.UserId, user.PositionId)
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

// GetUserDetail 获取用户详情的逻辑
func GetUserDetail(id int64) (data *models.ResUserDetail, err error) {
	return mysql.GetUserDetail(id)
}

// GetAllPerson 获取班级所有成员的逻辑
func GetAllPerson(id int64) (data []*models.ResUser, err error) {
	// 获取班级id
	classid, err := mysql.GetClassById(id)
	if err != nil {
		return nil, err
	}
	return mysql.GelAllPerson(classid)
}

// SearchPerson 查找用户的逻辑实现
func SearchPerson(s *models.ParamSearchUser) (data *models.ResUser, err error) {
	return mysql.SearchPerson(s)
}

// GetUserName 获取用户名称以及用户id的逻辑
func GetUserName(studentCard string, userId int64, positionId int64) (data *models.ResUerIdName, err error) {
	if positionId == 2 {
		return mysql.GetUserNameByStudentCard(studentCard)
	}
	classId, err := mysql.GetClassById(userId)
	if err != nil {
		return nil, err
	}
	return mysql.GetUserName(studentCard, classId)
}

// AddPerson 添加班级成员的接口
func AddPerson(s *models.ParamSearchUser, userId int64) error {
	// 获取被添加成员的班级Id 如果不为0 则无法被添加
	classId, err := mysql.GetClassById(s.UserId)
	if err != nil {
		return err
	}
	if classId != 0 {
		return ErrorInClass
	}
	// 获取班主任的班级id
	classId, err = mysql.GetClassById(userId)
	if err != nil {
		return err
	}
	if classId == 0 {
		return ErrorNotInClass
	}
	return mysql.ModifyClass(s, classId)
}

// RemovePerson 删除班级成员的接口
func RemovePerson(s *models.ParamSearchUser, userId int64) error {
	// 获取老师的班级id
	tClassId, err := mysql.GetClassById(userId)
	if err != nil {
		return err
	}
	if tClassId == 0 {
		return ErrorNotInClass
	}
	// 获取学生的班级id
	classId, err := mysql.GetClassById(s.UserId)
	if err != nil {
		return err
	}
	if classId == 0 {
		return ErrorNotInClass
	}
	// 判断二者是否在同一班级中
	if tClassId != classId {
		return ErrorNotSameClass
	}
	return mysql.ModifyClass(s, 0)
}

// UpdatePosition 修改职位
func UpdatePosition(p *models.ParamUpdatePosition,userId int64) error {
	// 获取老师的班级id
	tClassId, err := mysql.GetClassById(userId)
	if err != nil {
		return err
	}
	if tClassId == 0 {
		return ErrorNotInClass
	}
	// 获取学生的班级id
	classId, err := mysql.GetClassById(p.UserId)
	if err != nil {
		return err
	}
	if classId == 0 {
		return ErrorNotInClass
	}
	// 判断二者是否在同一班级中
	if tClassId != classId {
		return ErrorNotSameClass
	}
	// 判断这一职位在班级中是否已经存在了
	if p.PositionId!=1{
		has,err:=mysql.CheckPositionExist(tClassId,p.PositionId)
		if err!=nil{
			return err
		}
		if has{
			return ErrorPositionExist
		}
	}
	return mysql.UpdatePosition(p)
}

//-------------------------------------------------辅助函数-----------------------------------------

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
