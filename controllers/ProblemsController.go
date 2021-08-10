package controllers

import (
	"fmt"
	"strconv"
	"unioj/models"
)

type ProblemsController struct {
	BaseController
}

func (this *ProblemsController) GetPagesProblems() {
	pros, err := models.NewProblems().GetPagesProblems(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(*pros); i++ {
		models.O.LoadRelated(&((*pros)[i]), "ProblemType")
		(*pros)[i].Uid, _ = models.NewUser().GetUserByUid((*pros)[i].Uid.UId)
	}
	this.JsonResult(200, "问题列表加载成功", pros)
}

func (this *ProblemsController) GetProblemDetailById() {
	pid, _ := strconv.Atoi(this.Ctx.Input.Query("pid"))
	problems, err := models.NewProblems().GetProblemDetailById(pid)
	if err != nil {
		this.JsonResult(205, "服务器错误", err)
	}
	this.JsonResult(200, "ok", problems)
}

func (this *ProblemsController) GetProblemByTag() {
	tagname := this.Ctx.Input.Query("tagname")
	problems := models.NewProblems().GetProblemByTag(tagname)
	this.JsonResult(200, tagname, problems)
}
