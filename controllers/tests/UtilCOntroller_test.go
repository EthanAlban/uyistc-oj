package tests

import (
	"testing"
	"unioj/controllers"
)

func TestwetherToday(t *testing.T) {
	controllers.UtilsController.WetherToday(nil)
}