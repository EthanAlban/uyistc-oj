package models

type Language struct {
	Lid       int    `orm:"column(lid);pk"`
	Language_ string `orm:"column(language)"`
	Template  string `orm:"column(template)"`
}

// TableName 获取对应数据库表名.
func (l *Language) TableName() string {
	return "language"
}

// TableEngine 获取数据使用的引擎.
func (l *Language) TableEngine() string {
	return "INNODB"
}

func NewLanguage() *Language {
	return &Language{}
}
