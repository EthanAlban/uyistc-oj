package models

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