package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func sysinfoRouter() {
	beego.Router("api/sysinfo/get_sysinfo_list", &controllers.SysInfoController{}, "get:GetSysInfoList")
	beego.Router("api/sysinfo/set_info_read", &controllers.SysInfoController{}, "get:SetInfoRead")
}
