package tests

import (
	"github.com/wonderivan/logger"
	"testing"
	"unioj/models"
)

func TestGetAllAnnnocementsSortByTime(t *testing.T) {
	annocs := models.NewAnnonement().GetAllAnnnocementsSortByTime(10, 0)
	logger.Debug(annocs)
}
