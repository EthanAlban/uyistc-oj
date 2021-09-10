package healthy

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/wonderivan/logger"
	"io"
	"math"
	"net/http"
	"os"
	"sync"
	"time"
	"unioj-Judger/conf"
	"unioj-Judger/public"
	"unioj-Judger/service/mysql"
)

// RegisterJudger 判题器上线之后自动到数据库中进行注册
func RegisterJudger() {
	externalIp := os.Getenv("JUDGERIP")
	cfg := conf.CFG
	topics := cfg.Section("client").Key("kafka_topic").String()
	logger.Debug("注册判题器信息：externalIp", externalIp, "  topic: ", topics)
	// 检查当前有没有已经注册了的该ip的判题器
	nums := mysql.CheckJudgerExsit(externalIp)
	if nums == 0 {
		// 写入mysql
		mysql.RegisterJudgerIntoSql(externalIp, topics)
	} else {
		// 更新当前存在记录的判题器的在线状态和topic
		mysql.UpdateJudger(externalIp, topics)
	}

}

// DestroyJudger 销毁判题器
func DestroyJudger() {
	externalIp := os.Getenv("JUDGERIP")
	logger.Warn("正在销毁判题器：", externalIp)
	mysql.DestroyJudger(externalIp)
}

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
