package sql

import "time"

// TbActive 活动表
type TbActive struct {
	Id         int64
	Name       string	`xorm:"varchar(64)"`
	Creatos    string `xorm:"varchar(64)"`
	BeginTime  time.Time	`xorm:"notnull"`
	FinishTime time.Time `xorm:"notnull"`
	Category   int `xorm:"notnull"`
	Options    []string `xorm:"varchar(256)"`
	Base       `xorm:"extends"`
}
