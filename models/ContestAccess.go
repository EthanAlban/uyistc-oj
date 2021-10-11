package models

import (
	"fmt"
)

type ContestAccess struct {
	Cuid      string    `orm:"column(cuid);pk"`
	ContestId *Contests `orm:"column(contest_id);rel(fk)"`
	UserId    *User     `orm:"column(user_id);rel(fk)"`
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

// ContestAccess 返回对应用户是不使用权限参加竞赛
func (c *ContestAccess) ContestAccess(user User, contest_id string) bool {
	fmt.Println(user, contest_id)
	contest := NewContests().GetContestById(contest_id)
	var contestaccess ContestAccess
	err := O.QueryTable("contests_access").Filter("contest_id", contest).Filter("user_id", user).One(&contestaccess)
	if err != nil {
		return false
	}
	return true
}
