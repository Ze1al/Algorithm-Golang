package monotonic_stack

// 下一个更大元素
func nextGreaterElements(nums []int) []int {
	mergeNums := append(nums, nums...)

	record := make([]int, len(mergeNums))
	for i := 0; i < len(mergeNums); i++ {
		record[i] = -1
	}

	stack := []int{} // 存放的是下标

	for i := 0; i < len(mergeNums); i++ {
		top := len(stack) - 1
		for top >= 0 && mergeNums[stack[top]] <= mergeNums[i] {
			record[stack[top]] = mergeNums[i]
			stack = stack[:len(stack)-1]
			top -= 1
		}

		stack = append(stack, i)
	}

	res := make([]int, len(mergeNums))
	for i, v := range record {
		if v == -1 {
			res[i] = -1
			continue
		}
		res[i] = mergeNums[v]
	}

	return res[:len(nums)]
}
