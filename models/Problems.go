package models

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/wonderivan/logger"
	"strconv"
	"unioj/models/redisOP"
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
	TemplateC         string         `orm:"column(template_c)"`      //问题使用的初始标准代码
	TemplateGo        string         `orm:"column(template_go)"`     //问题使用的初始标准代码
	TemplateCplus     string         `orm:"column(template_c_plus)"` //问题使用的初始标准代码
	TemplateJava      string         `orm:"column(template_java)"`   //问题使用的初始标准代码
	TemplatePython    string         `orm:"column(template_python)"` //问题使用的初始标准代码

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
	proString, err := redisOP.RedisGetKey("problems_page_limit_" + strconv.Itoa(limit) + "_offset_" + strconv.Itoa(offset))
	if err != nil {
		_, err = O.QueryTable("problems").Limit(limit, offset).All(&pros)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		str, _ := json.Marshal(pros)
		redisOP.RedisSetKeyWithTimeout("problems_page_limit_"+strconv.Itoa(limit)+"_offset_"+strconv.Itoa(offset), string(str), 240)
		logger.Info("数据库加载问题列表")
		return &pros, err
	}
	json.Unmarshal([]byte(proString), &pros)
	logger.Info("缓存加载问题列表")
	return &pros, err
}

func (p *Problems) GetProblemDetailById(pid int) (*Problems, error) {
	var problem Problems
	err := O.QueryTable("problems").Filter("pid", pid).One(&problem)
	if err != nil {
		logger.Error("获取问题失败，id", pid, "  error:", err)
		return nil, err
	}
	return &problem, nil
}

func (p *Problems) GetProblemByTag(tagname string) *[]Problems {
	var problems []Problems
	O.QueryTable("problems").Filter("problem_type", NewTags().GetTagObjectByTagName(tagname)).All(&problems)
	return &problems
}

// IncProblemAcTimes 增加问题通过次数
func (p *Problems) IncProblemAcTimes(pid int) {
	qs := O.QueryTable("problems").Filter("pid", pid)
	_, err := qs.Update(orm.Params{
		"accept_submissions": orm.ColValue(orm.ColAdd, 1),
	})
	if err != nil {
		logger.Error("增加通过次数失败：", err)
	}
	var problem Problems
	err = qs.One(&problem)
	if err != nil {
		logger.Error(err)
	}
	acceptSubmissions := problem.AcceptSubmissions
	// 设置redis
	err = redisOP.RedisSetKeyWithTimeout("problem_"+strconv.Itoa(pid)+"_accept_submissions", strconv.Itoa(acceptSubmissions), 120)
	if err != nil {
		logger.Error("存储增加问题通过redis失败  ", err)
	}
	logger.Trace("增加问题通过" + strconv.Itoa(pid) + " 成功")
}

// IncProblemSubTimes 增加问题提交次数
func (p *Problems) IncProblemSubTimes(pid int) {
	qs := O.QueryTable("problems").Filter("pid", pid)
	_, err := qs.Update(orm.Params{
		"TotalSubmissions": orm.ColValue(orm.ColAdd, 1),
	})
	if err != nil {
		logger.Error("增加提交提交失败：", err)
	}
	var problem Problems
	qs.One(&problem)
	totalSubmissions := problem.TotalSubmissions
	// 设置redis
	err = redisOP.RedisSetKeyWithTimeout("problem_"+strconv.Itoa(pid)+"_total_submissions", strconv.Itoa(totalSubmissions), 120)
	if err != nil {
		logger.Error("存储增加问题提交redis失败  ", err)
	}
	logger.Trace("增加问题提交" + strconv.Itoa(pid) + " 成功")
}

// GetAcSubTimes 获取ac和提交次数
func (p *Problems) GetAcSubTimes(pid int) (int, int) {
	ac, err := redisOP.RedisGetKey("problem_" + strconv.Itoa(pid) + "_accept_submissions")
	sub, err := redisOP.RedisGetKey("problem_" + strconv.Itoa(pid) + "_total_submissions")
	if err != nil {
		logger.Error("获取ac_sub次数失败:", err)
		fmt.Println("--------------------------------------  ", pid)
		problem, _ := NewProblems().GetProblemDetailById(pid)
		logger.Debug("数据库加载sub_ac次数")
		err = redisOP.RedisSetKeyWithTimeout("problem_"+strconv.Itoa(pid)+"_total_submissions", strconv.Itoa(problem.TotalSubmissions), 120)
		err = redisOP.RedisSetKeyWithTimeout("problem_"+strconv.Itoa(pid)+"_accept_submissions", strconv.Itoa(problem.AcceptSubmissions), 120)
		return problem.AcceptSubmissions, problem.TotalSubmissions
	}
	logger.Trace("缓存加载sub_ac次数")
	ac_, _ := strconv.Atoi(ac)
	sub_, _ := strconv.Atoi(sub)
	return ac_, sub_
}

func (p *Problems) TotalProblems() int {
	//记录总的问题数量
	nums, err := redisOP.RedisGetKey("total_problems_count")
	if err != nil {
		// 查数据库
		count, err := O.QueryTable("problems").Count()
		if err != nil {
			logger.Error(err)
			return 0
		}
		redisOP.RedisSetKeyWithTimeout("total_problems_count", strconv.Itoa(int(count)), 240)
		logger.Info("查询数据库获取问题总数：", count)
		return int(count)
	}
	nums_, _ := strconv.Atoi(nums)
	logger.Info("缓存加载问题总数 ", nums_)
	return nums_
}
