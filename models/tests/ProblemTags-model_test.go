package tests

import (
	"testing"
	"unioj/models"
)

func TestGetProblemTagsById(t *testing.T) {
	models.NewProblemTags().GetProblemTagsById(1)
}
