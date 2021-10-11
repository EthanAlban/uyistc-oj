package models

import (
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
