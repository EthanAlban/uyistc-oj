package models

type ProblemTags struct {
	PTid int       `orm:"column(ptid);pk"`
	Tid  *Tags     `orm:"column(tid);rel(fk)"`
	Pid  *Problems `orm:"column(pid);rel(fk)"`
}

// TableName 获取对应数据库表名.
func (pts *ProblemTags) TableName() string {
	return "problem_tags"
}

// TableEngine 获取数据使用的引擎.
func (pts *ProblemTags) TableEngine() string {
	return "INNODB"
}

func NewProblemTags() *ProblemTags {
	return &ProblemTags{}
}

// GetProblemTagsById 按问题id查询问题所有的tagname
func (pts *ProblemTags) GetProblemTagsById(pid int) []string {
	tagsMap := NewTags().GetAllTagsMap()
	var problemTags []ProblemTags
	O.QueryTable("problem_tags").Filter("pid", pid).GroupBy("ptid").All(&problemTags)
	tags := make([]string, len(problemTags))
	for i := 0; i < len(problemTags); i++ {
		tags[i] = tagsMap[problemTags[i].Tid.Tid]
	}
	return tags
}
