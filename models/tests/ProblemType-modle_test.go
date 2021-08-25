package tests

import (
	"github.com/wonderivan/logger"
	"testing"
	"unioj/models"
)

func TestGetAllProblemTypes(t *testing.T) {
	pts := models.NewProblemType().GetAllProblemTypes()
	for i := 0; i < len(pts); i++ {
		logger.Debug(*(pts[i]))
	}
}
