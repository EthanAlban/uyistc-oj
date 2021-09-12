package tests

import (
	"fmt"
	"github.com/wonderivan/logger"
	"testing"
	"unioj/models"
)

func TestGetPagesProblems(t *testing.T) {
	pros, err := models.NewProblems().GetPagesProblems(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(*pros); i++ {
		models.O.LoadRelated(&((*pros)[i]), "Uid")
		models.O.LoadRelated(&((*pros)[i]), "ProblemType")
		logger.Debug((*pros)[i])
	}
}

func TestGetProblemByTag(t *testing.T) {
	logger.Debug(*(models.NewProblems().GetProblemByTag("快慢指针")))
}

func TestGetAcSubTimes(t *testing.T) {
	fmt.Println(models.NewProblems().GetAcSubTimes(1))
}
