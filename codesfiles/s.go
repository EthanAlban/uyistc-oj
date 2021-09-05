package codesfiles

func solveSudoku(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	// 一个二元组的切片  对应的位置表示空格出现位置的[i,j]
	var spaces [][2]int
	// 先用二重循环 把数独的基本情况记录下来
	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		// 如果已经算到最后一个空格了 证明算法结束
		if pos == len(spaces) {
			return true
		}
		// var spaces [][2]int
		// spaces[pos][0]: 第pos个空格的横坐标
		// spaces[pos][1]：第pos个空格的纵坐标
		i, j := spaces[pos][0], spaces[pos][1]
		// 从0-9遍历9个可能的填值
		for digit := byte(0); digit < 9; digit++ {
			// 满足数独要求  在一行一列一个九宫格都没有出现过
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				// 尝试把这个数放入数独
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1'
				// 如果这种填法 能够最终得到一个有效的解就是用这个数字
				if dfs(pos + 1) {
					return true
				}
				//否则当没有天国这个数字
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	dfs(0)
}
