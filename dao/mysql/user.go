package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"web-graduation/models"
	"web-graduation/models/sql"
)

const secret = "lxy1078014075"

// CheckUserExist 检查用户是否存在
func CheckUserExist(card string) (err error) {
	has, err := db.Where("student_card=?", card).Get(&sql.TbUser{})
	if err != nil {
		return err
	}
	if has {
		return ErrorUserExist
	}
	return
}

// InsertUser 新增用户
func InsertUser(u *sql.TbUser) (err error) {
	u.Password = encryptPassword(u.Password)
	u.PositionId = 1
	_, err = db.Cols("user_id", "user_name", "password", "student_card", "create_time", "update_time").Insert(u)
	return err
}

// LoginUser 用户登陆，数据库查找是否有该用户的数据
func LoginUser(u *sql.TbUser) (err error) {
	u.Password = encryptPassword(u.Password)
	has, err := db.Get(u)
	if err != nil {
		return err
	}
	// 判断用户是否存在 判断密码是否正确
	if !has {
		return ErrorInvalidPassword
	}
	return
}

// SetUpUser 用户信息的更新
func SetUpUser(u *sql.TbUser) (err error) {
	if u.Password != "" {
		u.Password = encryptPassword(u.Password)
	}
	affected, err := db.Where("user_id=?", u.UserId).Update(u)
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrorUserNotExist
	}
	return
}

// GetUserDetail 获取用户详情
func GetUserDetail(id int64) (data *models.ResUserDetail, err error) {
	fmt.Println(id)
	data = new(models.ResUserDetail)
	_, err = db.Table("tb_user").Join("inner", "tb_class", "tb_class.id=tb_user.class_id").
		Join("inner", "tb_position", "tb_position.id=tb_user.position_id").Where("user_id=?", id).Get(data)
	return
}

// GelAllPerson 获取班级所有成员
func GelAllPerson(classId int64) (data []*models.ResUser, err error) {
	data = make([]*models.ResUser, 0)
	err = db.Table("tb_user").Join("inner", "tb_position", "tb_position.id=tb_user.position_id").Where("class_id=?", classId).Find(&data)
	return
}

// SearchPerson 查找班级特定成员
func SearchPerson(s *models.ParamSearchUser) (data *models.ResUser, err error) {
	data = new(models.ResUser)
	_, err = db.Table("tb_user").Join("inner", "tb_position", "tb_position.id=tb_user.position_id").
		Where("user_id=?", s.UserId).Get(data)
	return
}

// GetUserName 获取用户名、用户id
func GetUserName(studentCard string, classId int64) (data *models.ResUerIdName, err error) {
	data = new(models.ResUerIdName)
	has, err := db.Table("tb_user").Cols("user_id", "user_name").Where("student_card=? and class_id=?", studentCard, classId).Get(data)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrorUserNotExist
	}
	return
}

// GetUserNameByStudentCard 通过学号获取用户名称、用户id
func GetUserNameByStudentCard(studentCard string) (data *models.ResUerIdName, err error) {
	data = new(models.ResUerIdName)
	has, err := db.Table("tb_user").Cols("user_id", "user_name").Where("student_card=?", studentCard).Get(data)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrorUserNotExist
	}
	return data, err
}

// GetClassByStudentCard 通过学号获取用户所在班级
func GetClassByStudentCard(studentCard string) (id int64, err error) {
	user := new(sql.TbUser)
	_, err = db.Cols("class_id").Where("student_card=?", studentCard).Get(user)
	if err != nil {
		return 0, err
	}
	return user.ClassId, nil
}

// ModifyClass 修改用户班级
func ModifyClass(s *models.ParamSearchUser, classId int64) error {
	u := new(sql.TbUser)
	u.ClassId = classId
	_, err := db.Where("user_id=?", s.UserId).Cols("class_id").Update(u)
	return err
}

// CheckPositionExist 检查职位是否存在
func CheckPositionExist(classId int64, positionId int64) (has bool, err error) {
	user := new(sql.TbUser)
	has, err = db.Where("class_id=? and position_id=?", classId, positionId).Get(user)
	return has, err
}

// UpdatePosition 修改职位
func UpdatePosition(p *models.ParamUpdatePosition) error {
	user := new(sql.TbUser)
	user.PositionId = p.PositionId
	_, err := db.Where("user_id=?", p.UserId).Cols("position_id").Update(user)
	return err
}

//-------------------------------------------------辅助函数-----------------------------------------

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
