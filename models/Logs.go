package models

import (
	uuid "github.com/satori/go.uuid"
	"github.com/wonderivan/logger"
	"strconv"
	"time"
)

type Logs struct {
	Lid     string    `orm:"column(lid);pk"`
	Type    int       `orm:"column(type)"`
	Uid     *User     `orm:"column(uid);rel(fk)"`
	Date    time.Time `orm:"column(date)"`
	Content string    `orm:"column(content)"`
}

// TableName 获取对应数据库表名.
func (l *Logs) TableName() string {
	return "logs"
}

// TableEngine 获取数据使用的引擎.
func (l *Logs) TableEngine() string {
	return "INNODB"
}

func NewLogs() *Logs {
	return &Logs{}
}

// AddLogging 添加日志
func (l *Logs) AddLogging(type_ string, userid *User, content string) bool {
	type__, err := strconv.Atoi(type_)
	if err != nil {
		logger.Error(err)
		return false
	}
	newLog := Logs{
		Type:    type__,
		Content: content,
		Date:    time.Now(),
		Lid:     (uuid.NewV4()).String(),
		Uid:     userid,
	}
	_, err = O.Insert(&newLog)
	if err != nil {
		logger.Error("插入日志失败:", err)
		return false
	}
	return true
}

// GetTypeLogsofUser 筛选特定用户的 特定类型的日志
func (l *Logs) GetTypeLogsofUser(type_ int, userid User) *[]Logs {
	var logs []Logs
	_, err := O.QueryTable("logs").Filter("Uid", userid).Filter("type", type_).All(&logs)
	if err != nil {
		logger.Error("查询日志失败", err)
		return nil
	}
	return &logs
}
