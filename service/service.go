package service

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"unioj-Judger/conf"
	"unioj-Judger/public"
	"unioj-Judger/server"
	"unioj-Judger/service/mysql"
	"unioj-Judger/service/structs"
)

// JudgeUserCode 接收用户代码进行评测
func JudgeUserCode(submission structs.Submission) {
	//logger.Debug("kafka读取submissoin: ", submission)
	server.JUDGER.Pid = submission.ProblemId.Pid
	server.JUDGER.Code = submission.Code
	server.JUDGER.Language = submission.Language.Language
	server.JUDGER.MemoryLimit = submission.ProblemId.MemoryLimit
	server.JUDGER.Language = submission.Language.Language
	server.JUDGER.Suffix = submission.Language.Suffix
	server.JUDGER.TimeLimit = time.Duration(submission.ProblemId.TimeLimit) * time.Millisecond
	//通过通道等待协程返回执行结果
	/*
		-2编译错误 -1答案错误 0正确 1计算超时  2超时 3内存超过 4运行时错误 5传送...  6判题中...  7部分正确
	*/
	res := make(chan public.JudgeResult, 1)
	go server.JUDGER.StartTask(res)
	Final := public.JudgeResult{
		Info: "",
		Code: -3,
	}
	for Final.Code == -3 {
		select {
		case Final = <-res:
			break
		}
	}
	ret := new(public.Ret)
	ret.Code = 200
	ret.Msg = "执行成功..."
	resMap := make(map[string]interface{})
	resMap["finalStat"] = Final.Code
	//替换掉敏感的服务器目录信息
	cfg := conf.CFG
	codesDir := cfg.Section("codesFile").Key("codesDir").String()
	Final.Info = strings.Replace(Final.Info, codesDir, "codesFile/", -1)
	resMap["info"] = Final.Info
	ret.Data = resMap
	//fmt.Println(ret)
	//将判题结果更新到数据库中
	submission.Result = Final.Code
	submission.ErrInfo = Final.Info
	submission.Score = Final.Score
	submission.LastTestcase = Final.LastTestcase
	submission.LastDesireOutput = Final.LastDesireOutput
	submission.LastOutput = Final.LastOutput
	mysql.UpdateSubmission(submission)
}

// GetParam 获取post参数
func GetParam(req *http.Request) map[string]string {
	// 根据 body 创建一个 json 解析器实例
	decoder := json.NewDecoder(req.Body)
	// 存放参数
	var params map[string]string
	// 解析参数 存入 map
	decoder.Decode(&params)
	return params
}
