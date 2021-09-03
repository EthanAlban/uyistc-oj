package codesfiles

import (
	"fmt"
	"testing"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// 二分查找向两边扩
	somePosition := binarySearch(nums, target)
	if somePosition == -1 {
		return []int{-1, -1}
	}
	left, right := findBoundary(nums, somePosition)
	return []int{left, right}
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2
	for left <= right {
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return -1
}

func findBoundary(nums []int, somePosition int) (int, int) {
	left, right := somePosition, somePosition
	target := nums[somePosition]
	if left-1 > 0 {
		for nums[left-1] == target {
			left--
			if left-1 < 0 {
				break
			}
		}
	}
	if right < len(nums)-1 {
		for nums[right+1] == target {
			right++
			if right+1 > len(nums)-1 {
				break
			}
		}
	}
	return left, right
}

func TestLeetcode(t *testing.T) {
	fmt.Println(searchRange([]int{1, 1, 2}, 1))
}
