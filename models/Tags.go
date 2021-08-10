package models

// Tags 保存系统中的问题标签
type Tags struct {
	Tid     int    `orm:"column(tid);pk"`
	TagName string `orm:"column(tag_name)"`
}

// TableName 获取对应数据库表名.
func (t *Tags) TableName() string {
	return "tags"
}

// TableEngine 获取数据使用的引擎.
func (t *Tags) TableEngine() string {
	return "INNODB"
}

func NewTags() *Tags {
	return &Tags{}
}
