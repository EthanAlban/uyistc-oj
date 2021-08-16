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
	"unioj/utils/captcha"
	"unioj/utils/kafka"
	"unioj/utils/session"
)

func main() {
	logger.Init()
	redisOP.Init()
	//注册session要用的结构体
	gob.Register(controllers.Weather{})
	gob.Register(models.User{})
	gob.Register(models.Problems{})
	session.InitSession()
	// 启动验证码服务
	go captcha.StartCaptchaServer()
	kafka.KafkaHealthCheck(beego.AppConfig.String("kafka_host"))
	port := conf.GetStringConfig("httpport")
	beego.Run("0.0.0.0:" + port)
	//orm.RunCommand()
}
