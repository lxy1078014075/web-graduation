package models

import "time"

// ResActiveList 活动列表的返还数据
type ResActiveList struct {
	ActiveId   int64  `json:"active_id"`
	ActiveName string `json:"active_name"`
	Creator    string `json:"creator"`
}

// ResActiveDetail 活动详情
type ResActiveDetail struct {
	ActiveName string    `json:"active_name"`
	Content    string    `json:"content"`
	FinishTime time.Time `json:"finish_time"`
}

// ResUserDetail 用户详情
type ResUserDetail struct {
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	StudentCard string `json:"student_card"`
	ClassName   string `json:"class_name"`
	Phone       string `json:"phone"`
	Position    string `json:"position" xorm:"'name'"`
}

// ResUsers 获取班级全部用户
type ResUsers struct {
	UserName string `json:"user_name"`
	StudentCard string `json:"student_card"`
	Position string `json:"position" xorm:"'name'"`
	Phone string `json:"phone"`
}
