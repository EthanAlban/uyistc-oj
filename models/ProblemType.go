package models

type ProblemType struct {
	Tid      int    `orm:"column(tid);pk"`
	TypeName string `orm:"column(type_name);size(20)"`
}

// TableName 获取对应数据库表名.
func (pt *ProblemType) TableName() string {
	return "problem_type"
}

// TableEngine 获取数据使用的引擎.
func (pt *ProblemType) TableEngine() string {
	return "INNODB"
}

func NewProblemType() *ProblemType {
	return &ProblemType{}
}

func (pt *ProblemType) GetAllProblemTypes() (pts []*ProblemType) {
	O.QueryTable("problem_type").All(&pts)
	return pts
}
