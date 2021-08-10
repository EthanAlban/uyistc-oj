package session

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	logger "unioj/logs"
)

var GlobalSessions *session.Manager

func InitSession() {
	address := beego.AppConfig.String("redis_host")
	sessionConfig := &session.ManagerConfig{
		CookieName:      "goo",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  address,
	}
	beego.BConfig.WebConfig.Session.SessionProviderConfig = address
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	var err error
	GlobalSessions, err = session.NewManager("redis", sessionConfig)
	if err != nil {
		logger.LogError("session初始化失败...,err:" + err.Error())
		fmt.Println("session初始化失败...,err:" + err.Error())
		return
	}
	go GlobalSessions.GC()
	fmt.Println("[3] INFO session初始化成功...")
}
