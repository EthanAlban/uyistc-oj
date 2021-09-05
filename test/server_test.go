package test

import (
	"sort"
	"testing"
	"unioj-Judger/server"
)

func TestProcess(t *testing.T) {
	server.Process()
}

func TestGetProcessInfo(t *testing.T) {
	server.GetProcessInfo("e0fe5683-3770-4b8b-aac8-8733ff6e6022")
}

func sdfg(nums []int) [][]int {
	var retNumsArr [][]int
	length := len(nums)
	if length < 3 {
		return retNumsArr
	}
	sort.Ints(nums)
	left, right := 0, length-1

	for i := 0; i < length; i++ {
		targetSum := 0 - nums[i]
		left = i + 1
		right = length - 1
		for left < right {
			if nums[left]+nums[right] == targetSum {
				intarr := []int{nums[i], nums[left], nums[right]}
				retNumsArr = append(retNumsArr, intarr)
			}
		}
	}
	return retNumsArr
}
