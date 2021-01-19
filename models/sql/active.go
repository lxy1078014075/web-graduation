package sql

import "time"

// TbActive 活动表
type TbActive struct {
	Id         int64
	ActiveId	int64	`json:"active_id" xorm:"notnull unique(idx_active_id)"`
	ActiveName       string `json:"active_name" xorm:"varchar(64)"`
	Creator    string `json:"creator" xorm:"varchar(64)"`
	ClassId    int64	`json:"class_id" xorm:"index(idx_class_id)"`
	Content  string `json:"content" xorm:"varchar(256) notnull"`
	FinishTime time.Time `json:"finish_time" xorm:"notnull"`
	Category   int       `json:"category" xorm:"notnull"`
	Options    []string  `json:"options" xorm:"varchar(256)"`
	Base       `xorm:"extends"`
}
