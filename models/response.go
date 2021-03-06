package models

import "time"

// ResActiveList 活动列表的返还数据
type ResActiveList struct {
	ActiveId   int64     `json:"active_id"`
	ActiveName string    `json:"active_name"`
	Creator    string    `json:"creator"`
	BeginTime  int64     `json:"begin_time"`
	CreateTime time.Time `json:"create_time"`
}

// ResActiveDetail 活动详情
type ResActiveDetail struct {
	ActiveId int64 `json:"active_id"`
	ActiveName string `json:"active_name"`
	Content    string `json:"content"`
	BeginTime  int64  `json:"begin_time"`
	FinishTime int64  `json:"finish_time"`
	Options []string `json:"options"`
	State      string `json:"state"`
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
	UserId   int64
	UserName string
}

// ResVoteDetail 获取 投票详情 和 投票结果 的返回值
type ResVoteDetail struct {
	Num int `json:"num"`
	UserName []string `json:"user_name" xorm:"'user_name'"`
}
