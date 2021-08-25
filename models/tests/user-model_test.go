package tests

import (
	"github.com/wonderivan/logger"
	"testing"
	"unioj/models"
)

func TestGetuserbyUsernamePwd(t *testing.T) {
	user := models.NewUser()
	user_, err := user.GetuserbyUsernamePwd("2814635354@qq.com", "741258", "email")
	if err == nil {
		logger.Debug(user_)
	}
	logger.Error(err)
}

func TestIsEmailUsed(t *testing.T) {
	logger.Debug(models.NewUser().IsEmailUsed("sad@qq.com"))
}
