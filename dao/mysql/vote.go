package mysql

import (
	"web-graduation/models"
	"web-graduation/models/sql"
)

// GetStudentCardById 通过 user_id 获取用户学号
func GetStudentCardById(userId int64) (studentCard string, err error) {
	_, err = db.Table("tb_user").Cols("student_card").Where("user_id=?", userId).Get(&studentCard)
	return
}

// Vote 投票
func Vote(p *models.ParamVote, studentCard string) error {
	vote := new(sql.TbVote)
	vote.State = 1
	vote.Result = p.Result
	affected, err := db.Where("active_id=? and student_card=?", p.ActiveId, studentCard).Update(vote)
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrorNotExist
	}
	return nil
}

// GetOptionsByActiveId 通过 active_id 获取选项
func GetOptionsByActiveId(activeId int64) ([]string, error) {
	active := new(sql.TbActive)
	_, err := db.Table("tb_active").Cols("options").Where("active_id=?", activeId).Get(active)
	return active.Options, err
}

// GetVoteDetailByOption
func GetVoteDetailByOption(activeId int64, option string, voteDetail *models.ResVoteDetail) error {
	err := db.Table("tb_vote").Join("left outer", "tb_user", "tb_vote.student_card=tb_user.student_card").
		Cols("user_name").Where("active_id=? and state=1 and result=?", activeId, option).Find(&voteDetail.UserName)
	return err
}

// GetNotVoteByActiveId 通过 active_id 获取没有参加投票的人
func GetNotVoteByActiveId(activeId int64) ([]string, error) {
	names := make([]string, 0)
	err := db.Table("tb_vote").Join("left outer", "tb_user", "tb_vote.student_card=tb_user.student_card").
		Cols("user_name").Where("active_id=? and state=0", activeId).Find(&names)
	return names, err
}
