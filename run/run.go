package main

import (
	"github.com/wonderivan/logger"
	"log"
	"net/http"
	"time"
	"unioj-Judger/conf"
	"unioj-Judger/service"
)

//从这里启动判题器
func main() {
	//初始化日志
	service.InitJudgerLogger()
	// 添加路由处理器
	http.HandleFunc("/healthy", service.HealthyCheck)
	http.HandleFunc("/usage_status", service.UsageStatus)
	//初始化配置文件
	conf_, err := conf.GetServerConfig()
	//初始化etcd连接
	timeoutSec, _ := conf.CFG.Section("etcd").Key("dialTimeout").Int()
	service.InitEtcd(time.Duration(int64(timeoutSec)) * time.Second)
	if err != nil {
		return
	}
	//初始化mysql  需要从beego的服务器拉取mysql的配置文件
	err = service.InitMysqlServer(conf_.Data.SqlUser, conf_.Data.SqlPassword, conf_.Data.SqlHost)
	if err != nil {
		return
	}
	//启动kafka的监听
	go service.NewKafkaConsumer(conf_.Data.KafkaHost)
	//启动定时删除过期文件任务
	cron := service.StartCronTask(120)
	defer cron.Stop()
	// 创建http服务端
	logger.Debug("judgeserver 启动成功...\n监听：0.0.0.0:7999")
	//fmt.Println("judgeserver 启动成功...\n监听：0.0.0.0:7999")
	//打印系统图形
	service.PrintSysLogo()
	err = http.ListenAndServe("0.0.0.0:7999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
