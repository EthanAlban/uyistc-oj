package models

//判题器模型

type Judger struct {
	Jid        int     `orm:"column(jid);pk"`
	JudgerHost string  `orm:"column(judger_host)"`
	MemUsage   float64 `orm:"column(mem_usage)"`
	LoadUsage  float64 `orm:"column(load_usage)"`
}
