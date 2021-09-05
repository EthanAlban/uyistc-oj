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

func isValidSudoku(board [][]byte) bool {
	// 九行的map
	LineMap := make([]map[byte]int, 9)
	for i := 0; i < 9; i++ {
		LineMap[i] = make(map[byte]int)
	}
	// 九列的map
	ColumeMap := make([]map[byte]int, 9)
	for i := 0; i < 9; i++ {
		ColumeMap[i] = make(map[byte]int)
	}
	// 九个方格的map
	SquareMap := make([]map[byte]int, 9)
	for i := 0; i < 9; i++ {
		SquareMap[i] = make(map[byte]int)
	}
	for i := 0; i < 9; i++ { // 行
		for j := 0; j < 9; j++ { // 列
			if board[i][j] != '.' {
				// 存入对应的行的map
				_, ok := LineMap[i][board[i][j]]
				if ok {
					return false
				} else {
					LineMap[i][board[i][j]] = j
				}
				// 存入对应的列的map
				_, ok = ColumeMap[j][board[i][j]]
				if ok {
					return false
				} else {
					ColumeMap[j][board[i][j]] = i
				}
				// 存入对应的3x3 map
				square := (i/3)*3 + (j / 3)
				_, ok = SquareMap[square][board[i][j]]
				if ok {
					return false
				} else {
					SquareMap[square][board[i][j]] = square
				}
			}
		}
	}
	return true
}

func TestLeetcode(t *testing.T) {
	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))
}
