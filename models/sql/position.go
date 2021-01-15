package sql

// TbPosition 用户身份,权限表
type TbPosition struct {
	Id   int64
	Name string `xorm:"varchar(64)"`
	Base `xorm:"extends"`
}
