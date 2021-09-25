package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func logsRouter() {
	beego.Router("api/logs/get_cur_user_loginhistory", &controllers.LogsController{}, "get:GetCurUserLoginHistory")
}
