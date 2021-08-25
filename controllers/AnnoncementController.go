package controllers

import (
	"github.com/wonderivan/logger"
	"unioj/models"
)

type AnnoncementController struct {
	BaseController
}

type limit_pages struct {
	limit  int
	offset int
}

func (this *AnnoncementController) GetAnnoncements() {
	var page limit_pages
	err := this.ParseForm(&page)
	if err != nil {
		logger.Error(500, "请同时给出limit,offset两个整形数字参数", err)
		//this.JsonResult(500, "请同时给出limit,offset两个整形数字参数", err)
		this.StopRun()
	}
	annonces := models.NewAnnonement().GetAllAnnnocementsSortByTime(page.limit, page.offset)
	for i := 0; i < len(*annonces); i++ {
		(*annonces)[i].Uid, _ = models.NewUser().GetUserByUid((*annonces)[i].Uid.UId)
	}
	this.JsonResult(200, "加载成功", annonces)
}
