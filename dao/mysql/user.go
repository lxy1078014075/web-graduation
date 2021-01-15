package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"web-graduation/models/sql"
)

const secret = "lxy1078014075"

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

func InsertUser(u *sql.TbUser) (err error) {
	password := encryptPassword(u.Password)
	u.Password = password
	_, err = db.Insert(u)
	return err
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
