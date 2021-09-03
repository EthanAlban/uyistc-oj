package service

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/wonderivan/logger"
	"io"
	"math"
	"net/http"
	"strings"
	"sync"
	"time"
	"unioj-Judger/conf"
	"unioj-Judger/public"
	"unioj-Judger/server"
)

// HealthyCheck judgeserver健康检查
func HealthyCheck(w http.ResponseWriter, req *http.Request) {
	ret := new(public.Ret)
	ret.Code = 200
	ret.Msg = "healthy"
	ret_json, _ := json.Marshal(ret)
	logger.Info("心跳检测--判题器工作正常")
	w.Header().Set("content-type", "text/json")
	io.WriteString(w, string(ret_json))
}

// UsageStatus 查看当前判题器所在的机器的各项参数指标 用于judgeserver的调度
func UsageStatus(w http.ResponseWriter, req *http.Request) {
	cfg := conf.CFG
	topics := cfg.Section("client").Key("kafka_topic").String()
	var wg sync.WaitGroup
	wg.Add(1)
	m, _ := mem.VirtualMemory()
	totalPercent, _ := cpu.Percent(1*time.Second, false)
	ret := public.SysUsage{
		MemLoad: math.Trunc(m.UsedPercent/0.01) * 0.01,
		LoadAvg: math.Trunc(totalPercent[0]/0.01) * 0.01,
		Topic:   topics,
	}
	logger.Info("内存统计:", ret.MemLoad, " 负载统计:", ret.LoadAvg, " topic:", topics)
	wg.Done()
	wg.Wait()
	b, _ := json.Marshal(ret)
	w.Header().Set("content-type", "text/json")
	io.WriteString(w, string(b))
}

// JudgeUserCode 接收用户代码进行评测
func JudgeUserCode(submission Submission) {
	//logger.Debug("kafka读取submissoin: ", submission)
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
	UpdateSubmission(submission)
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
