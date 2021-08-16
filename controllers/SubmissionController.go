package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"time"
	logger "unioj/logs"
	"unioj/models"
	"unioj/utils/kafka"
)

//用于调用判题器的controller
//用kafka的消息队列进行任务的分发

type SubmissionController struct {
	BaseController
}
type SubmisisonParams struct {
	Code      string `json:"code"`
	ProblemId int    `json:"problemId"`
	Language  string `json:"language"`
}

// SendTaskToKafka 发送到kafka队列中
func (this *SubmissionController) SendTaskToKafka() {
	userLogin := this.Ctx.Input.CruSession.Get("user_login")
	if userLogin == nil {
		this.JsonResult(204, "尚未登陆，不能进行提交", models.Submission{})
	}
	var taskParams SubmisisonParams
	body := this.Ctx.Input.RequestBody //这是获取到的json二进制数据
	err := json.Unmarshal(body, &taskParams)
	if err != nil {
		return
	}
	//解析二进制json，把结果放进ob中
	var task models.Submission
	lan := models.NewLanguage().GetLanguageById(taskParams.Language)
	task.Language = &lan
	problem, _ := models.NewProblems().GetProblemDetailById(taskParams.ProblemId)
	task.ProblemId = problem
	task.Code = taskParams.Code
	task.Result = -3
	task.Score = 0
	task.CreateTime = time.Now()
	submissionID := (uuid.NewV4()).String()
	task.SubmissionId = submissionID
	var user models.User
	user = userLogin.(models.User)
	task.UserName = user.UserName
	task.UserId = &user
	err = models.NewSubmission().InsertNewSubmission(task)
	if err != nil {
		logger.LogError("插入submission失败,请稍候重试... err:" + err.Error())
		this.JsonResult(205, "插入submission失败,请稍候重试...", err)
	}
	kafkaHost := beego.AppConfig.String("kafka_host")
	jugerTaskTopic := beego.AppConfig.String("juger_task_topic")
	err = kafka.SendToKafka(kafkaHost, jugerTaskTopic, task)
	if err != nil {
		this.JsonResult(205, "提交失败", err)
	}
	retMap := make(map[string]string)
	retMap["submissionID"] = submissionID
	this.JsonResult(200, "提交成功", retMap)
}
