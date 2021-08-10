package tests

import (
	"fmt"
	"testing"
	"unioj/models"
)

func TestGetuserbyUsernamePwd(t *testing.T) {
	user := models.NewUser()
	user_, err := user.GetuserbyUsernamePwd("2814635354@qq.com", "741258", "email")
	if err == nil {
		fmt.Printf("%v", user_)
	}
	fmt.Println(err)
}

func TestIsEmailUsed(t *testing.T) {
	fmt.Println(models.NewUser().IsEmailUsed("sad@qq.com"))
}
