package tests

import (
	"github.com/wonderivan/logger"
	"testing"
	"unioj/models"
)

func TestGetSysInfoList(t *testing.T) {
	logger.Debug(models.NewSysInfo().GetSysInfoList(10, 0))
}
