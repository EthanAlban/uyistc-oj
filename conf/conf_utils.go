package conf

import (
	"github.com/astaxie/beego/config"
	"github.com/wonderivan/logger"
)

func GetStringConfig(pathname string) string {
	conf, err := config.NewConfig("ini", "./conf/app.conf")
	if err != nil {
		logger.Error(err.Error())
		//fmt.Println(err.Error())
	}
	return conf.String(pathname)
}
