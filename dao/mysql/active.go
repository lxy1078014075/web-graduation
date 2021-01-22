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
	err = db.Table("tb_active").Where("class_id=?", classId).Find(&actives)
	fmt.Println(actives)
	return actives, err
}

// GetUserNameById 通过user_id获取用户名称
func GetUserNameById(userId int64) (creator string, err error) {
	user := new(sql.TbUser)
	_, err = db.Cols("user_name").Where("user_id=?", userId).Get(user)
	return user.UserName, err
}

// InsertActive 添加活动 使用事务进行添加，要为两张表同时添加
func InsertActive(active *sql.TbActive, studentCards []string) (err error) {
	session := db.NewSession()
	defer func() {
		if err != nil {
			session.Rollback()
		} else {
			err = session.Commit()
		}
		session.Close()
	}()
	err = session.Begin()

	_, err = session.Insert(active)
	if err != nil {
		return err
	}
	for _, studentCard := range studentCards {
		vote := new(sql.TbVote)
		vote.StudentCard = studentCard
		vote.ActiveId = active.ActiveId
		_, err = session.Omit("state", "result").Insert(vote)
		if err != nil {
			return err
		}
	}
	return err
}

// GetStudentCardsByClassId 通过班级获取所有的学号(教师除外)
func GetStudentCardsByClassId(classId int64) (studentCards []string, err error) {
	studentCards = make([]string, 0)
	err = db.Table("tb_user").Cols("student_card").Where("class_id=? and position_id!=?", classId, 2).Find(&studentCards)
	return
}

// GetActiveDetail 获取活动的详情
func GetActiveDetail(id int64) (data *models.ResActiveDetail, err error) {
	data = new(models.ResActiveDetail)
	_, err = db.Table("tb_active").Cols("active_id", "active_name", "content", "begin_time", "finish_time").Where("active_id=?", id).Get(data)
	return data, err
}

// GetCreatorByActiveId 通过 active_id 获取活动创建者
func GetCreatorByActiveId(id int64) (creator string, err error) {
	_, err = db.Table("tb_active").Cols("creator").Where("active_id=?", id).Get(&creator)
	return
}

// RemoveActive 删除活动 使用事务
func RemoveActive(activeId int64) (err error) {
	session := db.NewSession()
	defer func() {
		if err != nil {
			session.Rollback()
		} else {
			err = session.Commit()
		}
		session.Close()
	}()
	err = session.Begin()
	_, err = session.Where("active_id=?", activeId).Delete(&sql.TbActive{})
	if err != nil {
		return err
	}
	_, err = session.Where("active_id=?", activeId).Delete(&sql.TbVote{})
	if err != nil {
		return err
	}
	return
}
