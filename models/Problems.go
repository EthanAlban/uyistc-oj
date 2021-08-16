package models

import (
	"fmt"
	"strconv"
	logger "unioj/logs"
)

// Problems 记录所有问题的基本信息的表
type Problems struct {
	Pid               int            `orm:"column(pid);pk"`
	Title             string         `orm:"column(title);size(60)"`
	Level             int            `orm:"column(level)"`
	TotalSubmissions  int            `orm:"column(total_submissions)"`
	AcceptSubmissions int            `orm:"column(accept_submissions)"`
	Content           string         `orm:"column(content);"`
	TimeLimit         int            `orm:"column(time_limit)"`   // 运行时间限制
	MemoryLimit       int            `orm:"column(memory_limit)"` // 内存限制
	Hint              string         `orm:"column(hint)"`         // 题目提示
	ProblemType       *ProblemType   `orm:"column(problem_type);rel(fk)"`
	ProblemTags       []*ProblemTags `orm:"reverse(many)"`
	Uid               *User          `orm:"column(uid);rel(fk)"`
	InputDescription  string         `orm:"column(input_description)"`
	OutputDescription string         `orm:"column(output_description)"`
	Template          string         `orm:"column(template)"` //问题使用的初始标准代码
}

// TableName 获取对应数据库表名.
func (p *Problems) TableName() string {
	return "problems"
}

// TableEngine 获取数据使用的引擎.
func (p *Problems) TableEngine() string {
	return "INNODB"
}

func NewProblems() *Problems {
	return &Problems{}
}

func (p *Problems) GetPagesProblems(limit, offset int) (problems *[]Problems, err error) {
	var pros []Problems
	_, err = O.QueryTable("problems").Limit(limit, offset).All(&pros)
	if err != nil {
		return nil, err
	}
	return &pros, err
}

func (p *Problems) GetProblemDetailById(pid int) (*Problems, error) {
	var problem Problems
	err := O.QueryTable("problems").Filter("pid", pid).One(&problem)
	if err != nil {
		fmt.Println("获取问题失败，id", pid, "  error:", err)
		logger.LogError("获取问题失败，id" + strconv.Itoa(pid) + "  error:" + err.Error())
		return nil, err
	}
	return &problem, nil
}

func (p *Problems) GetProblemByTag(tagname string) *[]Problems {
	var problems []Problems
	O.QueryTable("problems").Filter("problem_type", NewTags().GetTagObjectByTagName(tagname)).All(&problems)
	return &problems
}
