package controllers

import (
	"fmt"
	"unioj/conf"
	"unioj/models"
)

type LogsController struct {
	BaseController
}

// GetCurUserLoginHistory 获取当前用户的所有的的了历史
func (this *LogsController) GetCurUserLoginHistory() {
	userLogin := this.Ctx.Input.CruSession.Get("user_login")
	if userLogin == nil {
		this.JsonResult(204, "尚未登陆无法查看")
	}
	loginHistory := models.NewLogs().GetTypeLogsofUser(conf.GetIntConfig("log_type_login"), userLogin.(models.User))
	fmt.Println(loginHistory)
	this.JsonResult(200, "登陆历史查询成功...", loginHistory)
}
