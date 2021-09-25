package models

import (
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
	"strconv"
	"time"
	"unioj/models/redisOP"
)

type Submission struct {
	SubmissionId     string    `orm:"column(submission_id);pk"`
	ProblemId        *Problems `orm:"column(problem_id);rel(fk)"`
	CreateTime       time.Time `orm:"column(create_time)"`
	UserId           *User     `orm:"column(user_id);rel(fk)"`
	UserName         string    `orm:"column(username)"`
	Code             string    `orm:"column(code)"`
	Result           int       `orm:"column(result)"`
	Language         *Language `orm:"column(language);rel(fk)"`
	Score            float64   `orm:"column(score)"`
	ErrInfo          string    `orm:"column(err_info)"`
	LastTestcase     string    `orm:"column(last_testcase)"`
	LastDesireOutput string    `orm:"column(last_desire_output)"`
	LastOutput       string    `orm:"column(last_output)"`
}

// TableName 获取对应数据库表名.
func (sb *Submission) TableName() string {
	return "submission"
}

// TableEngine 获取数据使用的引擎.
func (sb *Submission) TableEngine() string {
	return "INNODB"
}

func NewSubmission() *Submission {
	return &Submission{}
}

func (sb *Submission) InsertNewSubmission(sub Submission) error {
	_, err := O.Insert(&sub)
	if err != nil {
		return err
	}
	return err
}

// GetSubmissionByID 根据id查询submission
func (sb *Submission) GetSubmissionByID(submissionID string) (error, Submission) {
	var submission Submission
	err := O.QueryTable("submission").Filter("SubmissionId", submissionID).One(&submission)
	return err, submission
}

// GetProblemStatusLogin 通过UserId和问题id查询问题最后的最好的提交状态
func (sb *Submission) GetProblemStatusLogin(userId, pid int) int {
	var submissions []Submission
	_, err := O.QueryTable("submission").Filter("problem_id", pid).Filter("user_id", userId).OrderBy("create_time").All(&submissions)
	if err != nil {
		logger.Error("查询用户提交列表失败，err:", err)
	}
	if len(submissions) == 0 {
		return -3
	}
	for i := 0; i < len(submissions); i++ {
		if submissions[i].Result == 0 {
			return 0
		}
	}
	return submissions[0].Result
}

type GroupBy struct {
	Result string `json:"result"`
	Nums   int    `json:"nums"`
}

// GetSubmissionStaticForProblem 获取某个问题的通过统计信息
func (sb *Submission) GetSubmissionStaticForProblem(pid int) []GroupBy {
	value, err := redisOP.RedisGetKey(strconv.Itoa(pid) + "_submission_static")
	var subStastic []GroupBy
	if err != nil {
		// 查数据库
		_, err := O.Raw("SELECT result,count(*) as nums from submission WHERE problem_id=? GROUP BY result order by result", pid).QueryRows(&subStastic)
		if err != nil {
			logger.Error("查询提交统计信息失败,err:", err)
			return nil
		}
		// 存入redis
		bytes, _ := json.Marshal(subStastic)
		redisOP.RedisSetKeyWithTimeout(strconv.Itoa(pid)+"_submission_static", string(bytes), 120)
		logger.Debug("查询提交统计信息成功")
		return subStastic
	}
	json.Unmarshal([]byte(value), &subStastic)
	return subStastic
}

// GetUserSubmissions 分页查询某个用户的提交
func (sb *Submission) GetUserSubmissions(limit int, offset int, userid int) *[]Submission {
	var submissions []Submission
	O.QueryTable("submission").Filter("UserId", userid).Limit(limit, offset).OrderBy("-CreateTime").All(&submissions)
	return &submissions
}

// TotalSubmissions 获取某个用户的全部提交的数量
func (sb *Submission) TotalSubmissions(userid int) int {
	count, err := O.QueryTable("submission").Filter("UserId", userid).Count()
	if err != nil {
		logger.Error(err)
		return 0
	}
	return int(count)
}

// GetDoneProblems 获取某个用户做过的所有题目
func (sb *Submission) GetDoneProblems(userid int) (*[]Problems, map[int]int) {
	type problemCount struct {
		ProblemId int `json:"problem_id"`
	}
	var idlist []problemCount
	var problems []Problems
	status := make(map[int]int)
	O.Raw("SELECT problem_id FROM `submission`  WHERE user_id=? GROUP BY problem_id", userid).QueryRows(&idlist)
	fmt.Println(idlist)
	for _, id := range idlist {
		pro, err := NewProblems().GetProblemDetailById(id.ProblemId)
		if err != nil {
			return nil, nil
		}
		status[pro.Pid] = sb.GetProblemStatusLogin(userid, pro.Pid)
		problems = append(problems, *pro)
	}
	return &problems, status
}

// GetUserSubmissionProfile 得到用户提交总数  做过的题目总数 做过的所有问题
func (sb *Submission) GetUserSubmissionProfile(user User) (int, int, *[]Problems, map[int]int) {
	subCounts := sb.TotalSubmissions(user.UId)
	problems, status := sb.GetDoneProblems(user.UId)
	return subCounts, len(*problems), problems, status
}
