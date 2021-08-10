package conf

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	fmt.Println(GetStringConfig("dbhost"))
}
