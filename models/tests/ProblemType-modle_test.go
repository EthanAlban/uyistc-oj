package tests

import (
	"fmt"
	"testing"
	"unioj/models"
)

func TestGetAllProblemTypes(t *testing.T) {
	pts := models.NewProblemType().GetAllProblemTypes()
	for i := 0; i < len(pts); i++ {
		fmt.Println(*(pts[i]))
	}
}
