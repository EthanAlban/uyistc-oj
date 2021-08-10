package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func sysinfoRouter() {
	beego.Router("api/sysinfo/get_sysinfo_list", &controllers.SysInfoController{}, "get:GetSysInfoList")
}
