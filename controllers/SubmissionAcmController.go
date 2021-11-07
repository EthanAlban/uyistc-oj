package controllers

import (
	"github.com/wonderivan/logger"
	"math/rand"
	"strconv"
	"time"
	"unioj/conf"
	"unioj/models"
	"unioj/models/redisOP"
)

/*
	ACM赛制：每道题提交之后都有反馈，可以看到“通过”、“运行错误”、“答案错误”等等结果，
	但看不到错误的测试样例（leetcode周赛可以看到），每道题都有多个测试点，
	每道题必须通过了所有的测试点才算通过。每道题不限制提交次数，但没通过的话会有罚时，仅以最后一次提交为准。
	比赛过程中一般可以看到实时排名，通过题数相同的情况下按照答题时间+罚时来排名。
*/

// SubmissionAcmController 用于进行acm比赛的时候进行后端服务提供
type SubmissionAcmController struct {
	BaseController
	SubmissionController
}

//GetFinalInfoOfSubmission 查询提交的最终结果  只有在不是-3的时候才会进行返回
func (this *SubmissionAcmController) GetFinalInfoOfSubmission() {
	submissionID := this.Ctx.Input.Query("submissionID")
	contestID := this.Ctx.Input.Query("contestID")
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
	if submision.Result != 0 {
		// 如果是没通过的话就要加罚时
		penalty, _ := strconv.Atoi(conf.GetStringConfig("ACM_Penalty"))
		models.NewContestsAccess().AddPenalty(contestID, submision.UserId.UId, penalty)
		models.NewProblems().IncProblemAcTimes(submision.ProblemId.Pid)
	}
	submision.LastOutput = ""
	submision.LastDesireOutput = ""
	submision.LastOutput = ""
	this.JsonResult(200, "OK", submision)
}

//
