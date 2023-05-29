package monotonic_stack

// 单调栈解法
// 接雨水
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	stack := []int{}
	res := 0

	// 维护单调递减的栈
	stack = append(stack, 0)
	for i := 1; i < len(height); i++ {
		right := stack[len(stack)-1]
		if height[i] < height[right] {
			stack = append(stack, i)
		}
		if height[i] == height[right] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		}
		if height[i] > height[right] {
			for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
				bot := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if len(stack) != 0 {
					left := stack[len(stack)-1]

					height := min(height[left], height[i]) - height[bot]
					width := i - left - 1

					res += height * width
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}
