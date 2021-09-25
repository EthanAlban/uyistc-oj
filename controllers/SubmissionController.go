package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"github.com/wonderivan/logger"
	"math/rand"
	"strconv"
	"time"
	"unioj/models"
	"unioj/models/redisOP"
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
	lan := models.NewLanguage().GetLanguageByName(taskParams.Language)
	task.Language = &lan
	problem, _ := models.NewProblems().GetProblemDetailById(taskParams.ProblemId)
	task.ProblemId = problem
	task.Code = taskParams.Code
	//设置cookie到用户本地
	this.Ctx.SetCookie("judge_code", task.Code)
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
	// 增加提交次数
	models.NewProblems().IncProblemSubTimes(task.ProblemId.Pid)
	if err != nil {
		logger.Error("插入submission失败,请稍候重试... err:", err)
		//logger.LogError("插入submission失败,请稍候重试... err:" + err.Error())
		this.JsonResult(205, "插入submission失败,请稍候重试...", err)
	}
	kafkaHost := beego.AppConfig.String("kafka_host")
	//jugerTaskTopic := beego.AppConfig.String("juger_task_topic")
	//选择一个占用低的机器将任务发送到对应的topic中
	juger, err := models.NewJudger().GetLowUsageJudger()
	if err != nil {
		logger.Error(err)
		this.JsonResult(205, "无可用judger", err)
	}
	err = kafka.SendToKafka(kafkaHost, juger.KafkaTopic, task)
	if err != nil {
		this.JsonResult(205, "提交失败", err)
	}
	retMap := make(map[string]string)
	retMap["submissionID"] = submissionID
	this.JsonResult(200, "提交成功", retMap)
}

// IsSubmissionExsit 检查提交是否存在
func (this *SubmissionController) IsSubmissionExsit() {
	submissionID := this.Ctx.Input.Query("submissionID")
	err, _ := models.NewSubmission().GetSubmissionByID(submissionID)
	if err != nil && err.Error() == "<QuerySeter> no row found" {
		this.JsonResult(204, "没有对应的提交...")
		logger.Warn(err.Error())
		//logger.LogInfo(err.Error())
	}
	this.JsonResult(200, "ok")
}

// GetFinalInfoOfSubmission 查询提交的最终结果  只有在不是-3的时候才会进行返回
func (this *SubmissionController) GetFinalInfoOfSubmission() {
	submissionID := this.Ctx.Input.Query("submissionID")
	//避免在判题器不在线的时候或者长时间没有结果客户端一直询问 造成网络拥塞  记录查询次数
	// 用redis时间窗口限流 如果限流则会以50的概率不执行动作
	overflow, err := redisOP.RedisSetExpireLimit("limit_"+submissionID, 10, 5)
	if err == nil {
		if overflow {
			rand.Seed(time.Now().Unix())
			rate := rand.Intn(2)
			if rate == 0 {
				logger.Warn("用户请求被限流")
				this.JsonResult(203, "请求频繁请稍候重试", submissionID)
			}
		}
	}
	//正常逻辑
	err, submision := models.NewSubmission().GetSubmissionByID(submissionID)
	if err != nil {
		logger.Error(err.Error())
		if err.Error() == "<QuerySeter> no row found" {
			this.JsonResult(204, "没有对应的提交...")
			logger.Warn(err)
		} else {
			this.JsonResult(205, "服务器错误："+err.Error())
			logger.Error(err)
		}
	}
	if submision.Result == -3 {
		this.JsonResult(202, "判题中...")
	}
	if submision.Result == 0 {
		// 增加通过次数
		models.NewProblems().IncProblemAcTimes(submision.ProblemId.Pid)
	}
	this.JsonResult(200, "OK", submision)
}

// GetSubmissionStaticForProblem 获取某个问题的通过统计信息
func (this *SubmissionController) GetSubmissionStaticForProblem() {
	pid, _ := strconv.Atoi(this.Ctx.Input.Query("pid"))
	value := models.NewSubmission().GetSubmissionStaticForProblem(pid)
	this.JsonResult(200, "查询问题提交统计情况成功", value)
}

// GetUserSubmissions 分页获得当前用户的所有提交
func (this *SubmissionController) GetUserSubmissions() {
	limit, _ := strconv.Atoi(this.Ctx.Input.Query("limit"))
	offset, _ := strconv.Atoi(this.Ctx.Input.Query("offset"))
	userLogin := this.Ctx.Input.CruSession.Get("user_login")
	if userLogin != nil {
		userId := userLogin.(models.User).UId
		submissions := models.NewSubmission().GetUserSubmissions(limit, offset, userId)
		retmap := make(map[string]interface{})
		for i := 0; i < len(*submissions); i++ {
			models.O.LoadRelated(&((*submissions)[i]), "Language")
		}
		count := models.NewSubmission().TotalSubmissions(userId)
		retmap["total"] = count
		retmap["submissions"] = *submissions
		this.JsonResult(200, "查询提交列表成功", retmap)
	} else {
		this.JsonResult(204, "未登录无法查看提交记录")
	}
}

// GetUserSubmissionProfile 获得当前用的所有的提交数量
func (this *SubmissionController) GetUserSubmissionProfile() {
	userLogin := this.Ctx.Input.CruSession.Get("user_login")
	if userLogin == nil {
		this.JsonResult(204, "未登录无法查看提交记录")
	}
	retMap := make(map[string]interface{})
	retMap["sub_counts"], retMap["problems"], retMap["problems"], retMap["status"] = models.NewSubmission().GetUserSubmissionProfile(userLogin.(models.User))
	this.JsonResult(200, "查询用户做题详情成功...", retMap)
}
