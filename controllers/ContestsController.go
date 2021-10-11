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

// ContestAccess 查询当前用户是否有权限访问竞赛
func (this *ContestsController) ContestAccess() {
	contest_id := this.Ctx.Input.Query("contest_id")
	user, err := Userprofile(this.BaseController)
	if err != nil {
		this.JsonResult(204, "尚未登陆无法参加竞赛～")
	}
	res := models.NewContestsAccess().ContestAccess(user, contest_id)
	if res == true {
		this.JsonResult(200, "用户权限校验通过")
	} else {
		this.JsonResult(205, "用户权限校验失败")
	}
}

// ContestPasswordCheck 校验比赛密码
func (this *ContestsController) ContestPasswordCheck() {
	pwd := this.Ctx.Input.Query("pwd")
	contest_id := this.Ctx.Input.Query("contest_id")
	res := models.NewContests().ContestPasswordCheck(contest_id, pwd)
	if res {
		// 校验密码正确之后把对应用户加入要镇列表中
		user, err := Userprofile(this.BaseController)
		if err != nil {
			this.JsonResult(204, "尚未登陆无法参加竞赛～")
		}
		models.NewContests().AllowUserToContest(user, contest_id)
		this.JsonResult(200, "校验竞赛密码成功")
	}
	this.JsonResult(204, "校验竞赛密码失败")
}

// ContestProblemList 查询比赛中的所有的题目
func (this *ContestsController) ContestProblemList() {
	contest_id := this.Ctx.Input.Query("contest_id")
	problems, complete_num := models.NewContestProblems().ContestProblemList(contest_id)
	res := make(map[string]interface{})
	res["problems"] = problems
	res["complete_num"] = complete_num
	this.JsonResult(200, "查询竞赛问题列表成功", res)
}
