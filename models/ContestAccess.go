package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/wonderivan/logger"
	"time"
)

type ContestAccess struct {
	Cuid       string    `orm:"column(cuid);pk"`
	ContestId  *Contests `orm:"column(contest_id);rel(fk)"`
	UserId     *User     `orm:"column(user_id);rel(fk)"`
	EntryTime  time.Time `orm:"column(entry_time)"`
	FinishTime time.Time `orm:"column(finish_time)"`
	Penalty    float64   `orm:"column(penalty)"`
	Score      float64   `orm:"column(score)"`
}

// TableName 获取对应数据库表名.
func (c *ContestAccess) TableName() string {
	return "contests_access"
}

// TableEngine 获取数据使用的引擎.
func (c *ContestAccess) TableEngine() string {
	return "INNODB"
}

func NewContestsAccess() *ContestAccess {
	return &ContestAccess{}
}

// ContestAccess 返回对应用户是否有使用权限参加竞赛
func (c *ContestAccess) ContestAccess(user User, contest_id string) bool {
	contest := NewContests().GetContestById(contest_id)
	var contestaccess ContestAccess
	err := O.QueryTable("contests_access").Filter("contest_id", contest).Filter("user_id", user).One(&contestaccess)
	if err != nil {
		return false
	}
	return true
}

// AddPenalty 为某个竞赛的错误提交增加罚时
func (c *ContestAccess) AddPenalty(contest_id string, user_id string, penalty int) {
	contest := NewContests().GetContestById(contest_id)
	user, _ := NewUser().GetUserByUid(user_id)
	qs := O.QueryTable("contests_access").Filter("contest_id", contest).Filter("user_id", user)
	_, err := qs.Update(orm.Params{
		"accept_submissions": orm.ColValue(orm.ColAdd, penalty),
	})
	if err != nil {
		logger.Error("增加罚时失败：", err)
	}
}

// 查询用户在某个竞赛中的累积罚时
func GetTotalPenalty(userId, contestId string) {
	contest := NewContests().GetContestById(contest_id)

}
