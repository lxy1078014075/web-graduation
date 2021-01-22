package logic

import (
	"fmt"
	"time"
	"web-graduation/dao/mysql"
	"web-graduation/models"
	"web-graduation/models/sql"
	"web-graduation/pkg/snowflake"
)

// Active 获取活动列表
func Active(userId int64) (actives []*models.ResActiveList, err error) {
	// 获取用户班级信息
	classId, err := mysql.GetClassById(userId)
	if err != nil {
		return nil, err
	}
	if classId == 0 {
		return nil, ErrorNotInClass
	}
	// 获取活动列表
	actives, err = mysql.GetActiveList(classId)
	if err != nil {
		return nil, err
	}
	return
}

func GetActiveDetail(id int64) (data *models.ResActiveDetail, err error) {
	data, err = mysql.GetActiveDetail(id)
	if err != nil {
		return nil, err
	}
	// 获取活动状态
	now := time.Now().Unix()
	fmt.Println(data.BeginTime, now, data.FinishTime)
	switch {
	case now < data.BeginTime:
		data.State = "未开始"
	case data.BeginTime <= now && now < data.FinishTime:
		data.State = "进行中"
	case data.FinishTime >= now:
		data.State = "已结束"
	}
	return data, err
}

// AddActive 新建活动
func AddActive(param *models.ParamAddActive, userId int64) error {
	// 获取创建者以及班级信息
	creator, err := mysql.GetUserNameById(userId)
	if err != nil {
		return err
	}
	classId, err := mysql.GetClassById(userId)
	if err != nil {
		return err
	}
	if classId == 0 {
		return ErrorNotInClass
	}
	active := new(sql.TbActive)
	active.ActiveId = snowflake.GetID()
	active.ActiveName = param.Name
	active.Creator = creator
	active.ClassId = classId
	active.Content = param.Content
	active.BeginTime = param.BeginTime
	active.FinishTime = param.FinishTime
	active.Category = param.Category
	if active.Category == 1 {
		if param.Options == nil {
			return ErrorNeedOptions
		}
		active.Options = param.Options
	}
	// 获取班级成员的学号列表
	studentCards,err:=mysql.GetStudentCardsByClassId(classId)
	if err!=nil{
		return err
	}
	// 新建活动并为所有的班级成员分配活动
	return mysql.InsertActive(active,studentCards)
}

// RemoveActive 删除活动
func RemoveActive(activeId int64,userId int64,positionId int64) error {
	if positionId!=2{
		username,err:=mysql.GetUserNameById(userId)
		if err!=nil{
			return err
		}
		creator,err:=mysql.GetCreatorByActiveId(activeId)
		if err!=nil{
			return err
		}
		if creator==""{
			return ErrorActiveNotExist
		}
		if username!=creator{
			return ErrorNotSameCreator
		}
	}
	return mysql.RemoveActive(activeId)
}
