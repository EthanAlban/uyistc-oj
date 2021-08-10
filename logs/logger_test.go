package logs

import "testing"

func TestInit(t *testing.T) {
	Init()
	for i := 0; i < 11; i++ {
		LogError("test")
	}
}
