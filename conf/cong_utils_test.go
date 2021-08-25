package conf

import (
	"github.com/wonderivan/logger"
	"testing"
)

func TestGetConfig(t *testing.T) {
	logger.Debug(GetStringConfig("dbhost"))
	//fmt.Println(GetStringConfig("dbhost"))
}
