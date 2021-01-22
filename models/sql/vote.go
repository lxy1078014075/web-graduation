package sql

// TbVote 用户投票的表
type TbVote struct {
	Id          int64
	StudentCard string `xorm:"char(10) index(idx_studentCard_activeId)"`
	ActiveId    int64  `xorm:"index(idx_studentCard_activeId)"`
	State       int    `xorm:"notnull default(0)"`
	Result      string
	Base        `xorm:"extends"`
}
