package mysql

import (
	"crypto/md5"
	"encoding/hex"
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
	u.Identity=1
	_,err=db.Insert(u)
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

func SetUpUser(u *sql.TbUser) (err error) {
	affected, err := db.Where("user_id=?", u.UserId).Update(u)
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrorUserNotExist
	}
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
