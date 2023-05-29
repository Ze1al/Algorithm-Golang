package monotonic_stack

// 计算右边的温度，哪一天比今天高
// 依赖单调递增的栈，如果当天温度高于栈顶，则栈顶弹出，计算得到其下标值
func dailyTemperatures(temperatures []int) []int {
	stack := []int{} // 存放其下标

	res := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		stackIndex := len(stack) - 1
		// 如果栈顶小于当前输入值，则return
		for stackIndex >= 0 && temperatures[stack[stackIndex]] < temperatures[i] {
			val := stack[stackIndex]
			stack = stack[:len(stack)-1] // 弹出

			res[val] = i - val
			stackIndex -= 1
		}

		stack = append(stack, i)
	}
	return res
}
