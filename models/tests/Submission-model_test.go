package tests

import (
	"testing"
	"unioj/models"
)

func TestGetDoneProblems(t *testing.T) {
	models.NewSubmission().GetDoneProblems(1)
}
