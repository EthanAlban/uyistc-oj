package controllers

import (
	"strconv"
	"unioj/models"
)

type SysInfoController struct {
	BaseController
}

func (this *SysInfoController) GetSysInfoList() {
	var page limit_pages
	err := this.ParseForm(&page)
	if err != nil {
		this.JsonResult(500, "请同时给出limit,offset两个整形数字参数", err)
		this.StopRun()
	}
	sysinfos := models.NewSysInfo().GetSysInfoList(page.limit, page.offset)
	for i := 0; i < len(*sysinfos); i++ {
		(*sysinfos)[i].UId, _ = models.NewUser().GetUserByUid((*sysinfos)[i].UId.UId)
	}
	this.JsonResult(200, "", sysinfos)
}

// SetInfoRead 设置系统信息以读
func (this *SysInfoController) SetInfoRead() {
	Iid, _ := strconv.Atoi(this.Ctx.Input.Query("Iid"))
	models.NewSysInfo().SetSysInfoReadedById(Iid)
	res := make(map[string]bool)
	res["result"] = true
	this.JsonResult(200, "设置成功", res)
}
