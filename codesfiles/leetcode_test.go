package codesfiles

import (
	"fmt"
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	//  还是一个回溯的题目
	curentSolution := []int{}
	ret := [][]int{}
	ret2 := [][]int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ret = append(ret, append([]int(nil), curentSolution...))
			ret2 = append(ret2, curentSolution)
			fmt.Println("------------------------------")
			fmt.Println(curentSolution)
			fmt.Printf("ret: cap:%v len:%v val:%v\n", cap(ret), len(ret), ret)
			fmt.Printf("ret: cap:%v len:%v val:%v\n", cap(ret2), len(ret2), ret2)
			return
		}
		// 不采用当前的数字
		dfs(target, idx+1)
		// 用当前的数字
		if target-candidates[idx] >= 0 {
			curentSolution = append(curentSolution, candidates[idx])
			dfs(target-candidates[idx], idx)
			// 如果走下来了证明这条路走不通 需要把最后一个加进去的拿出来
			curentSolution = curentSolution[:len(curentSolution)-1]
		}
	}
	dfs(target, 0)
	return ret
}

func TestCombinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
