package session

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/wonderivan/logger"
)

var GlobalSessions *session.Manager

func InitSession() {
	address := beego.AppConfig.String("redis_host")
	sessionConfig := &session.ManagerConfig{
		CookieName:      "goo",
		EnableSetCookie: true,
		Gclifetime:      72000,
		Maxlifetime:     72000,
		Secure:          false,
		CookieLifeTime:  72000,
		ProviderConfig:  address,
	}
	beego.BConfig.WebConfig.Session.SessionProviderConfig = address
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	var err error
	GlobalSessions, err = session.NewManager("redis", sessionConfig)
	if err != nil {
		logger.Error("session初始化失败...,err:" + err.Error())
		return
	}
	go GlobalSessions.GC()
	logger.Debug("[3] INFO session初始化成功...")
}
