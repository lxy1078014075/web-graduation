package mysql

import (
	"fmt"
	"web-graduation/models"
	"web-graduation/models/sql"
)

// GetClassById 通过user_id获取用户所在班级
func GetClassById(userId int64) (id int64, err error) {
	user := new(sql.TbUser)
	_, err = db.Cols("class_id").Where("user_id=?", userId).Get(user)
	if err != nil {
		return 0, err
	}
	return user.ClassId, nil
}

// GetActiveList 获取活动列表
func GetActiveList(classId int64) (actives []*models.ResActiveList, err error) {
	actives = make([]*models.ResActiveList, 0)
	err = db.Table("tb_active").Cols("active_id", "active_name", "creator").Where("class_id=?", classId).Find(&actives)
	fmt.Println(actives)
	return actives, err
}

// GetUserNameById 通过user_id获取用户名称
func GetUserNameById(userId int64) (creator string, err error) {
	user := new(sql.TbUser)
	_, err = db.Cols("user_name").Where("user_id=?", userId).Get(user)
	return user.UserName, err
}



// InsertActive 添加活动
func InsertActive(active *sql.TbActive) error {
	_, err := db.Insert(active)
	return err
}

func GetActiveDetail(id int64) (data *models.ResActiveDetail,err error) {
	data=new(models.ResActiveDetail)
	_,err=db.Table("tb_active").Cols("active_name","content","finish_time").Get(data)
	return data,err
}