package mysql

import (
	"crypto/md5"
	"encoding/hex"
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
	data = new(models.ResUserDetail)
	_, err = db.Table("tb_user").Join("inner", "tb_class", "tb_class.id=tb_user.class_id").
		Join("inner", "tb_position", "tb_position.id=tb_user.position_id").Where("user_id=?", id).Get(data)
	return
}

// GelAllPerson 获取班级所有成员
func GelAllPerson(classId int64) (data []*models.ResUsers, err error) {
	data = make([]*models.ResUsers, 0)
	err = db.Table("tb_user").Join("inner", "tb_position", "tb_position.id=tb_user.position_id").Where("class_id=?", classId).Find(&data)

	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
