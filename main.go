package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"unioj/conf"
	"unioj/controllers"
	logger "unioj/logs"
	"unioj/models"
	"unioj/models/redisOP"
	_ "unioj/routers"
	"unioj/utils"
	"unioj/utils/captcha"
	"unioj/utils/kafka"
	"unioj/utils/session"
)

func main() {
	logger.InitJudgerLogger()
	redisOP.Init()
	//注册session要用的结构体
	gob.Register(controllers.Weather{})
	gob.Register(models.User{})
	gob.Register(models.Problems{})

	data := make(map[string]interface{})
	data["problem"] = models.Problems{}
	data["templates"] = models.Templates{}
	gob.Register(data)

	session.InitSession()
	// 启动验证码服务
	go captcha.StartCaptchaServer()
	kafka.KafkaHealthCheck(beego.AppConfig.String("kafka_host"))
	port := conf.GetStringConfig("httpport")
	//定时 检测judger的健康状态
	interval, _ := beego.AppConfig.Int("healthy_check_interval")
	cron := utils.StartJudgerHealtyCheck(interval)
	defer cron.Stop()
	utils.CheckJudgerHealth()
	utils.PrintSysLogo()
	beego.Run("0.0.0.0:" + port)
	//orm.RunCommand()
}
