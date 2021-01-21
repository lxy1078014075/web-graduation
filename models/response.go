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
	Age         int    `json:"age"`
}

// ResUser 获取班级用户的信息
type ResUser struct {
	UserName    string `json:"user_name"`
	StudentCard string `json:"student_card"`
	Position    string `json:"position" xorm:"'name'"`
	Phone       string `json:"phone"`
}

// ResUerIdName 获取用户Id 和 用户名
type ResUerIdName struct {
	UserId int64
	UserName string
}
