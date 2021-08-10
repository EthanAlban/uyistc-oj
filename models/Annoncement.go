package models

import "time"

type Annnocement struct {
	Cid     int       `orm:"column(cid);pk"`
	Uid     *User     `orm:"column(uid);rel(fk)"`
	Content string    `orm:"column(content);size(255)"`
	Time    time.Time `orm:"column(time)"`
	Title   string    `orm:"column(title)"`
	Cover   string    `orm:"column(cover);default('')"`
}

// TableName 获取对应数据库表名.
func (an *Annnocement) TableName() string {
	return "annocement"
}

// TableEngine 获取数据使用的引擎.
func (an *Annnocement) TableEngine() string {
	return "INNODB"
}

func NewAnnonement() *Annnocement {
	return &Annnocement{}
}

func (an *Annnocement) GetAllAnnnocementsSortByTime(limit, offset int) (annocements *[]Annnocement) {
	var annonces []Annnocement
	O.QueryTable("annocement").OrderBy("time").Limit(limit, offset).All(&annonces)
	return &annonces
}
