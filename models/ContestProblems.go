package models

type ContestProblems struct {
	Cpid        string    `orm:"column(cpid);pk"`
	ContestId   *Contests `orm:"column(contest_id);rel(fk)"`
	ProblemId   *Problems `orm:"column(problem_id);rel(fk)"`
	CompleteNum int       `orm:"column(complete_num)"`
}

// TableName 获取对应数据库表名.
func (cp *ContestProblems) TableName() string {
	return "contest_problems"
}

// TableEngine ContestProblems.
func (cp *ContestProblems) TableEngine() string {
	return "INNODB"
}

func NewContestProblems() *ContestProblems {
	return &ContestProblems{}
}

// ContestProblemList 查询比赛中的所有的题目
func (cp *ContestProblems) ContestProblemList(contest_id string) (*[]Problems, *map[int]int) {
	contest := NewContests().GetContestById(contest_id)
	var contest_problems []ContestProblems
	completeNum := make(map[int]int)
	O.QueryTable("contest_problems").Filter("contest_id", contest).All(&contest_problems)
	for i := 0; i < len(contest_problems); i++ {
		O.LoadRelated(&(contest_problems[i]), "problem_id")
		// 查询该竞赛中这个问题的完成次数
		completeNum[contest_problems[i].ProblemId.Pid] = contest_problems[i].CompleteNum
	}
	var problems []Problems
	for i := 0; i < len(contest_problems); i++ {
		problems = append(problems, *contest_problems[i].ProblemId)
	}
	return &problems, &completeNum
}
