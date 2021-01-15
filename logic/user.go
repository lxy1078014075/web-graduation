package logic

import (
	"web-graduation/dao/mysql"
	"web-graduation/models"
	"web-graduation/models/sql"
	"web-graduation/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if err = mysql.CheckUserExist(p.StudentCard); err != nil {
		return err
	}
	// 生成ID
	userId:=snowflake.GetID()
	user:=&sql.TbUser{
		UserId:      userId,
		UserName:    p.Username,
		Password:    p.Password,
		StudentCard: p.StudentCard,
	}
	return mysql.InsertUser(user)
}
