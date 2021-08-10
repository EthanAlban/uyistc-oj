package conf

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) { //sad
	fmt.Println(GetStringConfig("dbhost"))
}
