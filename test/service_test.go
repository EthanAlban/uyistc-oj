package test

import (
	"fmt"
	"testing"
	"unioj-Judger/conf"
	"unioj-Judger/service"
)

func TestUsage(t *testing.T) {
	//service.UsageStatus()
}

func TestConfig(t *testing.T) {
	config, err := conf.GetServerConfig()
	if err != nil {
		return
	}
	fmt.Println(*config)
}

func TestConsumer(t *testing.T) {
	//service.Consumer()
}

func TestProducer(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

func TestPrintLogo(t *testing.T) {
	service.PrintSysLogo()
}

func TestKafkaConsumer(t *testing.T) {
	service.Consumer("39.105.152.206:9092")
}

func TestKafkaConsumer2(t *testing.T) {
	service.Consumer("39.105.152.206:9092")
}
