package monotonic_stack

// 接雨水
// https://leetcode.cn/problems/trapping-rain-water/

// 双指针解法
// 按照列计算
func trap1(height []int) int {
	var res int
	for i := 0; i < len(height); i++ {
		if i == 0 || i == len(height)-1 {
			continue
		}

		leftHeight, rightHeight := height[i], height[i]
		for r := i + 1; r < len(height); r++ {
			if rightHeight < height[r] {
				rightHeight = height[r]
			}
		}

		for l := i - 1; l >= 0; l-- {
			if leftHeight < height[l] {
				leftHeight = height[l]
			}
		}

		diff := min(leftHeight, rightHeight) - height[i]
		if diff > 0 {
			res += diff * 1
		}
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
