package sql

// TbVote 用户投票的表
type TbVote struct {
	Id          int64
	ActiveId    int64  `xorm:"index(idx_activeId_studentCard)"`
	StudentCard string `xorm:"char(10) index(idx_activeId_studentCard)"`
	State       int    `xorm:"notnull default(0)"`
	Result      string
	Base        `xorm:"extends"`
}
