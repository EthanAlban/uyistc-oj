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

// GetAllTags 获得系统中的所有tag
func (t *Tags) GetAllTags() (*[]Tags, error) {
	var tags []Tags
	_, err := O.QueryTable("tags").All(&tags)
	if err != nil {
		return nil, err
	}
	return &tags, nil
}

// GetTagObjectByTagName 根据tagname返回tag对象
func (t *Tags) GetTagObjectByTagName(tagname string) *Tags {
	var tag Tags
	O.QueryTable("tags").Filter("tag_name", tagname).One(&tag)
	return &tag
}
