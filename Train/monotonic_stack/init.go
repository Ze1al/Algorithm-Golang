package monotonic_stack

// 单调栈，寻找一个元素的右边或者左边第一个比自己大或者小的元素位置，则使用单调栈

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
