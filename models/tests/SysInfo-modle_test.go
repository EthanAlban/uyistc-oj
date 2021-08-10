package tests

import (
	"fmt"
	"testing"
	"unioj/models"
)

func TestGetSysInfoList(t *testing.T) {
	fmt.Println(models.NewSysInfo().GetSysInfoList(10, 0))
}
