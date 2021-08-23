package models

import (
	"encoding/json"
	"unioj/models/redisOP"
)

type Language struct {
	Lid       int    `orm:"column(lid);pk"`
	Language_ string `orm:"column(language)"`
	Template  string `orm:"column(template)"`
}

// Templates 序列化问题中的template字段以支持多种语言形成一个数组
type Templates struct {
	LanName string `json:"lan_name"`
	Code    string `json:"code"`
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

// GetAllLanges 获取所有的语言名称及其模板代码
func (l *Language) GetAllLanges() (*[]Language, error) {
	var lans []Language
	//读redis看有没有
	lans_str, err := redisOP.RedisGetKey("sysLanguages")
	if err == nil {
		json.Unmarshal([]byte(lans_str), &lans)
		return &lans, err
	} else {
		_, err = O.QueryTable("language").All(&lans)
		if err != nil {
			return nil, err
		}
		return &lans, err
	}
}

func (l *Language) GetLanguageById(languageName string) Language {
	var lan Language
	O.QueryTable("language").Filter("Language_", languageName).One(&lan)
	return lan
}
