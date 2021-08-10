package logs

import (
	"fmt"
	_ "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	confutils "unioj/conf"
)

var uniLogger *logs.BeeLogger

func Init() {
	logFile := confutils.GetStringConfig("beego_log_location")
	uniLogger = logs.NewLogger(10)                                   // 创建一个日志记录器，参数为缓冲区的大小
	err := uniLogger.SetLogger("file", `{"filename":"`+logFile+`"}`) // 设置日志记录方式：本地文件记录
	if err != nil {
		uniLogger.Error("err:", err)
		return
	}
	uniLogger.SetLevel(logs.LevelDebug) // 设置日志写入缓冲区的等级
	uniLogger.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）
	uniLogger.Flush()                   // 将日志从缓冲区读出，写入到文件
	fmt.Println("[1] INFO unilogger初始化成功...")
	//uniLogger.Close()
}

func LogError(err_msg string) {
	uniLogger.Error(err_msg)
}

func LogInfo(info string) {
	uniLogger.Info(info)
}
