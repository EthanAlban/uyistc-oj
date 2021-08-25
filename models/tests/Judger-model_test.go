package tests

import (
	"fmt"
	"testing"
	"unioj/models"
)

func TestCheckJudgerHealth(t *testing.T) {
	fmt.Println(models.NewJudger().GetAllJudger())
}
