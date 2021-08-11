package tests

import (
	"testing"
	"unioj/models"
)

func TestGetAllTagsMap(t *testing.T) {
	models.NewTags().GetAllTagsMap()
}
