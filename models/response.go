package models

import "time"

// ResActiveList 活动列表的返还数据
type ResActiveList struct {
	ActiveId   int64  `json:"active_id"`
	ActiveName string `json:"active_name"`
	Creator    string `json:"creator"`
}

type ResActiveDetail struct {
	ActiveName string `json:"active_name"`
	Content string `json:"content"`
	FinishTime time.Time `json:"finish_time"`
}
