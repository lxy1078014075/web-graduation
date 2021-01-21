package sql

// TbUser 用户信息表
type TbUser struct {
	Id          int64
	UserId      int64  `xorm:"notnull unique(idx_user_id)"`
	UserName    string `xorm:"varchar(64) notnull"`
	Password    string `xorm:"notnull"`
	Age         int
	Email       string
	Gender      string `xorm:"char(3) notnull default('男') "`
	StudentCard string `xorm:"char(10) unique(idx_student_card)"`
	ClassId     int64  `xorm:"index(idx_class_id)"`
	Phone       string `xorm:"char(11)"`
	Residence   string `xorm:"varchar(64)"`
	PositionId  int64  `xorm:"default(1)"`
	Base        `xorm:"extends"`
}
