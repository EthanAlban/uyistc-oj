package logs

import (
	"github.com/wonderivan/logger"
)

func InitJudgerLogger() {
	err := logger.SetLogger("./conf/logger.json")
	if err != nil {
		logger.Error("Judgerloger初始化失败:" + err.Error())
		return
	}
	logger.Debug("Judgerloger初始化成功...")
	return
}
