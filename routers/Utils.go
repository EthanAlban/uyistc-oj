package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func utilsRouter() {
	beego.Router("api/utils/weather", &controllers.UtilsController{}, "get:WetherToday")
	beego.Router("api/utils/heartbeat", &controllers.UtilsController{}, "get:Heartbeat")
	beego.Router("api/utils/captcha", &controllers.UtilsController{}, "get:Captcha")
	beego.Router("api/utils/set_info_read", &controllers.UtilsController{}, "get:SetInfoRead")
}
