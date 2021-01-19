package sql

type TbScore struct {
	Id                int64
	StudentCard       string  `xorm:"char(10) unique(idx_student_card)"`
	BaseScore         float64 //基础分
	GPA               float64 //绩点
	Sport             float64 //体育分
	CourseCompetition float64 //课程设计鼻塞分
	Art               float64 //
	GroupActive       float64
	TotalScore        float64
	State             int //状态
}
