package sql

type TbZonece struct {
	Id int64
	StudentCard string `xorm:"char(10) unique(idx_student_card)"`
	ActiveId int64
	CourseCompetitionId int64
	ArtId int64
	GroupActiveId int64
	State int
}
