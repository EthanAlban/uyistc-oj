package controllers

import "unioj/models"

type LanguageController struct {
	BaseController
}

func (this *LanguageController) GetAllLanges() {
	lans, err := models.NewLanguage().GetAllLanges()
	if err != nil {
		this.JsonResult(200, "语言模板加载成功", lans)
	}
	this.JsonResult(205, "server error")
}
