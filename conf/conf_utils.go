package conf

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func GetStringConfig(pathname string) string {
	conf, err := config.NewConfig("ini", "/home/et/Desktop/Projects/beego/unioj/conf/app.conf")
	if err != nil {
		fmt.Println(err.Error())
	}
	return conf.String(pathname)
}
