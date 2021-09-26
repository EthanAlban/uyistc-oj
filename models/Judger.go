package models

import (
	"github.com/wonderivan/logger"
)

//判题器模型

type Judger struct {
	Jid        int     `orm:"column(jid);pk"`
	JudgerHost string  `orm:"column(judger_host)"`
	MemUsage   float64 `orm:"column(mem_usage)"`
	LoadUsage  float64 `orm:"column(load_usage)"`
	Online     int     `orm:"column(online)"`
	KafkaTopic string  `orm:"column(kafka_topic)"`
}

// TableName 获取对应数据库表名.
func (jg *Judger) TableName() string {
	return "judger"
}

// TableEngine 获取数据使用的引擎.
func (jg *Judger) TableEngine() string {
	return "INNODB"
}

func NewJudger() *Judger {
	return &Judger{}
}

func (j *Judger) GetAllJudger() *[]Judger {
	var judgers []Judger
	_, err := O.QueryTable("judger").All(&judgers)
	if err != nil {
		logger.Error(err)
	}
	return &judgers
}

// UpdateJudgerStatus 更新judger的状态
func (j *Judger) UpdateJudgerStatus(host string, online int, mem_usage, load_usage float64, topic string) error {
	judger := Judger{JudgerHost: host}
	err := O.Read(&judger, "judger_host")
	if err == nil {
		judger.Online = online
		judger.MemUsage = mem_usage
		judger.LoadUsage = load_usage
		//judger.LoadUsage = load_usage
		judger.KafkaTopic = topic
		if _, err := O.Update(&judger); err == nil {
			logger.Debug("更新 " + host + " 状态成功...")
		} else {
			logger.Error("更新 " + host + " 状态失败...")
			logger.Error(err)
		}
	} else {
		logger.Error(err)
	}
	return err
}

// GetLowUsageJudger 选择一个cpu负载低的机器进行发送kafka任务
func (j *Judger) GetLowUsageJudger() (*Judger, error) {
	var judger Judger
	err := O.QueryTable("judger").Filter("online", 1).OrderBy("load_usage").One(&judger)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return &judger, nil
}
