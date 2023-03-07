package monotonic_stack

// 每日温度
// https://leetcode.cn/problems/daily-temperatures/
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	// 维护一个从左到右是递减关系的栈，如果出现比栈头大的元素，则将栈头元素弹出，直到找到比他大的元素为止
	stack := []int{}
	stack = append(stack, 0)

	for i := 1; i < len(temperatures); i++ {
		right := stack[len(stack)-1]

		if temperatures[i] < temperatures[right] {
			stack = append(stack, i)
		}
		if temperatures[i] == temperatures[right] {
			stack = append(stack, i)
		}
		if temperatures[i] > temperatures[right] {
			for len(stack) > 0 {
				tempRight := stack[len(stack)-1]
				if temperatures[tempRight] < temperatures[i] {
					stack = stack[:len(stack)-1]
					res[tempRight] = i - tempRight
				} else {
					break
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}
