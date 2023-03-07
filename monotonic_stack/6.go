package monotonic_stack

// 柱状图中最大的矩形
// https://leetcode.cn/problems/largest-rectangle-in-histogram/

// 双指针方法
func largestRectangleArea1(heights []int) int {
	var res int
	for i := 0; i < len(heights); i++ {
		left := i
		for left >= 0 {
			if heights[left] < heights[i] {
				break
			}
			left--
		}

		right := i
		for right <= len(heights)-1 {
			if heights[right] < heights[i] {
				break
			}
			right++
		}

		area := (right - left - 1) * heights[i]
		res = max(res, area)
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
