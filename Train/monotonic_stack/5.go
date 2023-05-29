package monotonic_stack

// 柱状图中最大的矩形
// 单调栈
// 从左到右是递增的
// 出现小于栈顶的元素
func largestRectangleArea(heights []int) int {
	maxArea := 0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	stack := []int{0} // 栈
	for i := 1; i < len(heights); i++ {
		top := stack[len(stack)-1]

		if heights[i] > heights[top] {
			stack = append(stack, i)
		} else if heights[i] == heights[top] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && heights[i] < heights[top] {
				bot := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if len(stack) != 0 {
					leftIndex := stack[len(stack)-1]
					rightIndex := i

					h := heights[bot]
					w := rightIndex - leftIndex - 1
					maxArea = max(h*w, maxArea)
				}
			}
			stack = append(stack, i)
		}
	}

	return maxArea
}
