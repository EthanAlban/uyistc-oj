package controllers

import (
	"strconv"
	"unioj/models"
)

type ContestsController struct {
	BaseController
}

// GetAllContests 分页查询竞赛列表
func (this *ContestsController) GetAllContests() {
	limit, _ := strconv.Atoi(this.Ctx.Input.Query("limit"))
	offset, _ := strconv.Atoi(this.Ctx.Input.Query("offset"))
	contests := models.NewContests().GetALlContest(limit, offset)
	this.JsonResult(200, "获取比赛列表成功...", *contests)
}

// GetContestQuantity 查询竞赛的总数
func (this *ContestsController) GetContestQuantity() {
	this.JsonResult(200, "查询竞赛数量成功", models.NewContests().GetContestQuantity())
}
