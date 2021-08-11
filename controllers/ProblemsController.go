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
	//先读session
	currentProblemDetail := this.Ctx.Input.CruSession.Get("current_problem_detail")
	if currentProblemDetail != nil {
		if (currentProblemDetail).(models.Problems).Pid == pid {
			this.JsonResult(200, "缓存加载当前问题", currentProblemDetail)
		}
	}
	problem, err := models.NewProblems().GetProblemDetailById(pid)
	if err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			this.JsonResult(204, "未找到相应id的问题", err)
		}
		this.JsonResult(205, "服务器错误", err)
	}
	problem.Uid, _ = models.NewUser().GetUserByUid(problem.Uid.UId)
	//存入session
	if this.Ctx.Input.CruSession == nil {
		this.StartSession()
	}
	this.Ctx.Input.CruSession.Set("current_problem_detail", *problem)

	////设置cookie每次尝试从cookie加载当前题目内容
	//problemByte,_ := json.Marshal(problem)
	//// 第三个参数是时间 单位是秒 ，如果不设置时间，Cookie只在本次回话有效，默认存活时间为3600秒
	//this.Ctx.SetCookie("current_problem_detail",string(problemByte))
	this.JsonResult(200, "加载问题成功", problem)
}

func (this *ProblemsController) GetProblemByTag() {
	tagname := this.Ctx.Input.Query("tagname")
	problems := models.NewProblems().GetProblemByTag(tagname)
	this.JsonResult(200, tagname, problems)
}

// GetProblemTagsById 按问题id查出该问题的tags
func (this *ProblemsController) GetProblemTagsById() {
	pid, _ := strconv.Atoi(this.Ctx.Input.Query("pid"))
	tags := models.NewProblemTags().GetProblemTagsById(pid)
	this.JsonResult(200, "tags加载完成", tags)
}
