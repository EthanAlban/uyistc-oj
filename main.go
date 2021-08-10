package main

import (
	"github.com/astaxie/beego"
	"unioj/conf"
	logger "unioj/logs"
	"unioj/models/redisOP"
	_ "unioj/routers"
	"unioj/utils/captcha"
	"unioj/utils/session"
)

func main() {
	logger.Init()
	redisOP.Init()
	session.InitSession()
	// 启动验证码服务
	go captcha.StartCaptchaServer()
	port := conf.GetStringConfig("httpport")
	beego.Run("0.0.0.0:" + port)
	//orm.RunCommand()
}
