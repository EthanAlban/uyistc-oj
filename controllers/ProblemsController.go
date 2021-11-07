package controllers

import (
	"fmt"
	"github.com/wonderivan/logger"
	"strconv"
	"unioj/models"
)

type ProblemsController struct {
	BaseController
}

// GetPagesProblems 分页查询题目列表  得到列表的同时通过用户id查询当前用户是否已经通过这个题目
func (this *ProblemsController) GetPagesProblems() {
	limit, _ := strconv.Atoi(this.Ctx.Input.Query("limit"))
	offset, _ := strconv.Atoi(this.Ctx.Input.Query("offset"))
	fmt.Println("limit:", limit, " offset:", offset)
	// 获取当前的登陆用户
	var userId string
	userLogin := this.Ctx.Input.CruSession.Get("user_login")
	if userLogin != nil {
		userId = userLogin.(models.User).UId
	} else {
		userId = "-1"
	}
	pros, err := models.NewProblems().GetPagesProblems(limit, offset)
	if err != nil {
		logger.Error(err)
		//fmt.Println(err)
	}
	// map[问题id]问题状态
	problemStatus := make(map[int]int)
	for i := 0; i < len(*pros); i++ {
		models.O.LoadRelated(&((*pros)[i]), "ProblemType")
		(*pros)[i].Uid, _ = models.NewUser().GetUserByUid((*pros)[i].Uid.UId)
		// 加载缓存用户做题信息
		problemStatus[(*pros)[i].Pid] = models.NewSubmission().GetProblemStatusLogin(userId, (*pros)[i].Pid)

	}
	retmap := make(map[string]interface{})
	retmap["results"] = problemStatus
	retmap["problems"] = pros
	retmap["total"] = models.NewProblems().TotalProblems()
	this.JsonResult(200, "问题列表加载成功", retmap)
}

// GetProblemDetailById 根据问题id查询单个问题的详情
func (this *ProblemsController) GetProblemDetailById() {
	pid, _ := strconv.Atoi(this.Ctx.Input.Query("pid"))
	//先读session
	currentProblemDetail := this.Ctx.Input.CruSession.Get("current_problem_detail")
	if currentProblemDetail != nil {
		if (currentProblemDetail).(map[string]interface{})["problem"].(models.Problems).Pid == pid {
			this.JsonResult(200, "缓存加载当前问题", currentProblemDetail)
		}
	}
	problem, err := models.NewProblems().GetProblemDetailById(pid)
	templates := make(map[string]string)
	if problem.TemplateC != "" {
		templates["C"] = problem.TemplateC
	}
	if problem.TemplateGo != "" {
		templates["Golang"] = problem.TemplateGo
	}
	if problem.TemplateCplus != "" {
		templates["C++"] = problem.TemplateCplus
	}
	if problem.TemplateJava != "" {
		templates["Java"] = problem.TemplateJava
	}
	if problem.TemplatePython != "" {
		templates["Python"] = problem.TemplatePython
	}
	if err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			this.JsonResult(204, "未找到相应id的问题", err)
		}
		this.JsonResult(205, "服务器错误", err)
	}
	problem.Uid, _ = models.NewUser().GetUserByUid(problem.Uid.UId)
	data := make(map[string]interface{})
	data["problem"] = problem
	data["templates"] = templates
	//存入session
	if this.Ctx.Input.CruSession == nil {
		this.StartSession()
	}
	this.Ctx.Input.CruSession.Set("current_problem_detail", data)

	////设置cookie每次尝试从cookie加载当前题目内容
	//problemByte,_ := json.Marshal(problem)
	//// 第三个参数是时间 单位是秒 ，如果不设置时间，Cookie只在本次回话有效，默认存活时间为3600秒
	//this.Ctx.SetCookie("current_problem_detail",string(problemByte))
	this.JsonResult(200, "加载问题成功", data)
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

// GetProblemAcSubTimes  获取对应问题的通过次数和提交次数
func (this *ProblemsController) GetProblemAcSubTimes() {
	pid, _ := strconv.Atoi(this.Ctx.Input.Query("pid"))
	ac, sub := models.NewProblems().GetAcSubTimes(pid)
	retMap := make(map[string]int)
	retMap["ac"] = ac
	retMap["submissions"] = sub
	this.JsonResult(200, "获取问题提交次数成功", retMap)
}

// GetProblemSupportLanguage 获取某个问题支持使用的问题列表
func (this *ProblemsController) GetProblemSupportLanguage() {
	pid, _ := strconv.Atoi(this.Ctx.Input.Query("pid"))
	lans := models.NewProblems().GetProblemSupportLanguage(pid)
	this.JsonResult(200, "查询问题接受语言成功", lans)
}

// GetAcAndPenalty 获取比赛中问题的通过次数总罚时
func (this *ProblemsController) GetAcAndPenalty() {
	var userId string
	userLogin := this.Ctx.Input.CruSession.Get("user_login")
	if userLogin != nil {
		userId = userLogin.(models.User).UId
	} else {
		userId = "-1"
	}
	contest_id := this.Ctx.Input.Query("contest_id")
	pid, _ := strconv.Atoi(this.Ctx.Input.Query("pid"))
	actime, penalty := models.NewContestProblems().GetAcAndPenalty(contest_id, userId, pid)
	resMap := make(map[string]interface{})
	resMap["actimes"] = actime
	resMap["penalty"] = penalty
	this.JsonResult(200, "查询总罚时和通过次数成功", resMap)
}
