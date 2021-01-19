package sql

// TbUser 用户信息表
type TbUser struct {
	Id          int64
	UserId      int64  `xorm:"notnull unique(idx_user_id)"`
	UserName    string `xorm:"varchar(64) notnull"`
	Password    string `xorm:"notnull"`
	Email       string
	Gender      int    `xorm:"notnull default(1) "`
	StudentCard string `xorm:"char(10) unique(idx_student_card)"`
	ClassId     int64    `xorm:"index(idx_class_id)"`
	Phone       string `xorm:"char(11)"`
	Identity    int `xorm:"default(1)"`
	Base        `xorm:"extends"`
}
