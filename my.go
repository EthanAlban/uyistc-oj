package unioj_judger

func nextPermutation(nums []int) {
	n := len(nums)
	//找到第一个不是从右到左递增的数字考虑让他和仅仅比他大一点的数字进行交换
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	//	从左到右找到第一个比nums[i]大的数字
	if i >= 0 {
		j := n - 1
		for j > 0 && nums[i] >= nums[j] {
			j--
		}
		//	交换i,j位置的数字
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 由于此时[i+1:]是递减的，我们需要让他更小所以要让[i+1:]翻转一下
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
