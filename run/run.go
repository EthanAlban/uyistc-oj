package main

import (
	"github.com/wonderivan/logger"
	"log"
	"net/http"
	"time"
	"unioj-Judger/conf"
	"unioj-Judger/service"
	"unioj-Judger/service/healthy"
	"unioj-Judger/service/mysql"
	"unioj-Judger/service/utils"
)

//从这里启动判题器
func main() {
	//初始化日志
	service.InitJudgerLogger()
	// 添加路由处理器
	http.HandleFunc("/healthy", healthy.HealthyCheck)
	http.HandleFunc("/usage_status", healthy.UsageStatus)
	//初始化配置文件
	conf_, err := conf.GetServerConfig()
	//初始化mysql  需要从beego的服务器拉取mysql的配置文件
	err = mysql.InitMysqlServer(conf_.Data.SqlUser, conf_.Data.SqlPassword, conf_.Data.SqlHost)
	if err != nil {
		return
	}

	//初始化etcd连接
	timeoutSec, _ := conf.CFG.Section("etcd").Key("dialTimeout").Int()
	service.InitEtcd(time.Duration(int64(timeoutSec)) * time.Second)
	if err != nil {
		return
	}
	//启动kafka的监听
	go service.NewKafkaConsumer(conf_.Data.KafkaHost)
	//启动定时删除过期文件任务
	cron := service.StartCronTask(120)
	//fmt.Println("judgeserver 启动成功...\n监听：0.0.0.0:7999")
	//注册判题器
	healthy.RegisterJudger()
	//打印系统图形
	utils.PrintSysLogo()
	go func() {
		for {
			select {
			case <-service.DeathKafka:
				// 清理任务 判题器挂了要删除对应记录
				logger.Error("kafka挂了")
				cron.Stop()
				healthy.DestroyJudger()
				logger.Fatal("判题器撤出...")
			}
		}
	}()
	//logger.Fatal("kafka挂了")
	// 创建http服务端
	logger.Debug("judgeserver 启动成功...   监听：0.0.0.0:7999")
	err = http.ListenAndServe("0.0.0.0:7999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
