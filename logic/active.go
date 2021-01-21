package logic

import (
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
	if classId==0{
		return nil, ErrorNotInClass
	}
	// 获取活动列表
	actives, err = mysql.GetActiveList(classId)
	if err != nil {
		return nil, err
	}
	return
}

func GetActiveDetail(id int64) (data *models.ResActiveDetail,err error) {
	return mysql.GetActiveDetail(id)
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
	if classId==0{
		return ErrorNotInClass
	}
	active := new(sql.TbActive)
	active.ActiveId = snowflake.GetID()
	active.ActiveName = param.Name
	active.Creator = creator
	active.ClassId = classId
	active.Content = param.Content
	active.FinishTime = time.Unix(param.FinishTime, 0)
	active.Category = param.Category
	if active.Category == 1 {
		if param.Options == nil {
			return ErrorNeedOptions
		}
		active.Options = param.Options
	}
	return mysql.InsertActive(active)
}
