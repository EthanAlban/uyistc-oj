package tests

import (
	"fmt"
	"testing"
	"unioj/models"
)

func TestGetAllAnnnocementsSortByTime(t *testing.T) {
	annocs := models.NewAnnonement().GetAllAnnnocementsSortByTime(10, 0)
	fmt.Printf("%v", annocs)
}
