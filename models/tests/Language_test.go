package tests

import (
	"fmt"
	"testing"
	"unioj/models"
)

func TestGetLanguageByName(t *testing.T) {
	fmt.Println(models.NewLanguage().GetLanguageByName("Golang"))
}
