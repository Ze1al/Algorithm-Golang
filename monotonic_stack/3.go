package monotonic_stack

// 下一个更大的元素
// https://leetcode.cn/problems/next-greater-element-ii/
func nextGreaterElements(nums []int) []int {
	mergeNum := append(nums, nums...)

	res := make([]int, len(mergeNum))
	for i := 0; i < len(mergeNum); i++ {
		res[i] = -1
	}

	// 维持单调递减的栈，从左到右变小的栈
	stack := []int{}
	stack = append(stack, 0)

	for i := 1; i < len(mergeNum); i++ {
		right := stack[len(stack)-1]
		if mergeNum[right] > mergeNum[i] {
			stack = append(stack, i)
		}
		if mergeNum[right] == mergeNum[i] {
			stack = append(stack, i)
		}

		if mergeNum[right] < mergeNum[i] {
			for len(stack) > 0 && mergeNum[stack[len(stack)-1]] < mergeNum[i] {
				temp := stack[len(stack)-1]
				res[temp] = mergeNum[i]

				stack = stack[:len(stack)-1]
			}

			stack = append(stack, i)
		}
	}

	return res[:len(nums)]
}
