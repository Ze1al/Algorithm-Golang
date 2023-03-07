package monotonic_stack

// 最大矩形
func largestRectangleArea(heights []int) int {
	var (
		stack []int
		res   int
	)
	newHeight := make([]int, len(heights)+2)
	for i := 1; i < len(heights)+1; i++ {
		newHeight[i] = heights[i-1]
	}

	stack = append(stack, 0)

	for i := 1; i < len(newHeight); i++ {
		right := stack[len(stack)-1]
		if newHeight[right] < newHeight[i] {
			stack = append(stack, i)
		} else if newHeight[right] == newHeight[i] {
			stack = append(stack, i)
		} else {
			// 如果栈中一直小于输入元素，则一直往外弹
			for len(stack) > 0 && newHeight[stack[len(stack)-1]] > newHeight[i] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					left := stack[len(stack)-1]
					right := i
					w := right - left - 1
					h := newHeight[top]
					res = max(res, w*h)
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}
