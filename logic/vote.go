package logic

import (
	"web-graduation/dao/mysql"
	"web-graduation/models"
)

// Vote 用户投票的逻辑实现
func Vote(p *models.ParamVote, userId int64) error {
	studentCard, err := mysql.GetStudentCardById(userId)
	if err != nil {
		return err
	}
	return mysql.Vote(p, studentCard)
}

// VoteDetail 获取投票详情的逻辑实现
func VoteDetail(activeId int64) (data []*models.ResVoteDetail, err error) {
	// 获取活动的选项数量
	options, err := mysql.GetOptionsByActiveId(activeId)
	if err != nil {
		return
	}
	if len(options) == 0 {
		return nil, ErrorActiveHasProblem
	}
	data = make([]*models.ResVoteDetail, len(options))
	for i, option := range options {
		voteDetail := new(models.ResVoteDetail)
		err = mysql.GetVoteDetailByOption(activeId, option, voteDetail)
		if err != nil {
			return nil, err
		}
		voteDetail.Num = len(voteDetail.UserName)
		data[i]=voteDetail
	}
	return
}

func VoteResult(activeId int64) (data []*models.ResVoteDetail,err error) {
	// 获取活动的选项数量
	options, err := mysql.GetOptionsByActiveId(activeId)
	if err != nil {
		return
	}
	if len(options) == 0 {
		return nil, ErrorActiveHasProblem
	}
	data = make([]*models.ResVoteDetail, len(options))
	for i, option := range options {
		voteDetail := new(models.ResVoteDetail)
		err = mysql.GetVoteDetailByOption(activeId, option, voteDetail)
		if err != nil {
			return nil, err
		}
		if i==0{	//默认选项 没有选的人都划分到这一选项中
			names,err:=mysql.GetNotVoteByActiveId(activeId)
			if err!=nil{
				return nil,err
			}
			voteDetail.UserName=append(voteDetail.UserName,names...)
		}
		voteDetail.Num = len(voteDetail.UserName)
		data[i]=voteDetail
	}
	return
}
