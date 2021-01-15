package sql

// TbClass 班级表
type TbClass struct {
	Id        int64
	ClassName string `xorm:"notnull"`
	Base      `xorm:"extends"`
}