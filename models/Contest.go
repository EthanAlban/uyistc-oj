package models

import (
	uuid "github.com/satori/go.uuid"
	"github.com/wonderivan/logger"
	"strconv"
	"time"
	"unioj/models/redisOP"
)

type Contests struct {
	Cid         string    `orm:"column(cid);pk"`
	ContestName string    `orm:"column(contest_name)"`
	ContestDesc string    `orm:"column(contest_desc)"`
	ContestType int       `orm:"column(contest_type)"`
	Creater     *User     `orm:"column(creater);rel(fk)"`
	StartTime   time.Time `orm:"column(start_time)"`
	EndingTime  time.Time `orm:"column(ending_time)"`
	Password    string    `orm:"column(password)"`
}

// TableName 获取对应数据库表名.
func (c *Contests) TableName() string {
	return "contests"
}

// TableEngine 获取数据使用的引擎.
func (c *Contests) TableEngine() string {
	return "INNODB"
}

func NewContests() *Contests {
	return &Contests{}
}

// GetALlContest 分页查询竞赛列表
func (c *Contests) GetALlContest(limit, offset int) *[]Contests {
	var contests []Contests
	_, err := O.QueryTable("contests").Limit(limit, offset).All(&contests)
	if err != nil {
		logger.Error("获取所有比赛列表失败... ", err)
		return nil
	}
	for i := 0; i < len(contests); i++ {
		contests[i].Creater, _ = NewUser().GetUserByUid(contests[i].Creater.UId)
	}
	return &contests
}

// GetContestQuantity 获取竞赛的数量
func (c *Contests) GetContestQuantity() int {
	key, err := redisOP.RedisGetKey("contest_all_number")
	if err != nil {
		// 需要查数据库
		var contests []Contests
		all, _ := O.QueryTable("contests").All(&contests)
		redisOP.RedisSetKeyWithTimeout("contest_all_number", strconv.Itoa(int(all)), 240)
		return int(all)
	}
	val, err := strconv.Atoi(key)
	if err != nil {
		return 0
	}
	return val
}

// GetContestById 按照id查询竞赛
func (c *Contests) GetContestById(contest_id string) *Contests {
	var contest Contests
	O.QueryTable("contests").Filter("cid", contest_id).One(&contest)
	return &contest
}

func (c *Contests) ContestPasswordCheck(id string, pwd string) bool {
	var contest Contests
	O.QueryTable("contests").Filter("cid", id).One(&contest)
	if contest.Password == pwd {
		return true
	}
	return false
}

func (c *Contests) AllowUserToContest(user User, id string) {
	contest := NewContests().GetContestById(id)
	var contestAccess ContestAccess
	contestAccess.UserId = &user
	contestAccess.ContestId = contest
	contestAccess.Cuid = (uuid.NewV4()).String()
	O.Insert(&contestAccess)
}
