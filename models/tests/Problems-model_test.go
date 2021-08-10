package tests

import (
	"fmt"
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
		fmt.Println((*pros)[i])
	}
}
