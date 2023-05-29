package monotonic_stack

// 接雨水
// 方法一：中心扩散法
// 按照行计算，容易出现重复的值
// 按照列计算比较简单,找到列左右两边的最大值
func trap1(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	sum := 0
	for i := 1; i < len(height)-1; i++ {
		leftHeight, rightHeight := height[i], height[i]
		for left := i - 1; left >= 0; left-- {
			if height[left] > leftHeight {
				leftHeight = height[left]
			}
		}

		for right := i + 1; right <= len(height)-1; right++ {
			if height[right] > rightHeight {
				rightHeight = height[right]
			}
		}

		if min(leftHeight, rightHeight) < height[i] {
			continue
		}
		area := min(leftHeight, rightHeight) - height[i]

		sum += area
	}

	return sum
}

// 方法二：维护两个数组，记录从左到右的最大值，记录从右到左的最大值
func trap2(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	sum := 0

	maxLeft := make([]int, len(height))  // 从左到右，递增数组
	maxRight := make([]int, len(height)) // 从右到左，递增数组

	var maxLeftNum int
	for i := 0; i < len(height); i++ {
		if maxLeftNum < height[i] {
			maxLeftNum = height[i]
		}
		maxLeft[i] = maxLeftNum
	}

	var maxRightNum int
	for i := len(height) - 1; i > 0; i-- {
		if maxRightNum < height[i] {
			maxRightNum = height[i]
		}
		maxRight[i] = maxRightNum
	}

	for i := 1; i < len(height)-1; i++ {
		if maxLeft[i] == height[i] || maxRight[i] == height[i] {
			continue
		}

		area := min(maxLeft[i], maxRight[i]) - height[i]
		sum += area
	}

	return sum
}

// 使用单调栈
// stack 从栈顶到栈底为从小到大的顺序，记录栈中的下标值
// [5,4,1] 1为底部，4为左边，添加的元素为右边
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	sum := 0

	stack := []int{}
	stack = append(stack, 0)
	for i := 1; i < len(height); i++ {
		// 如果小于栈顶，则添加
		if height[stack[len(stack)-1]] > height[i] {
			stack = append(stack, i)
			continue
		}
		// 如果等于栈顶，则替换
		if height[stack[len(stack)-1]] == height[i] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
			continue
		}

		// 如果大于栈顶元素，则可以计算值
		if height[stack[len(stack)-1]] < height[i] {
			// 本质是一行一行计算的，所以需要while循环
			for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
				buttonIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if len(stack) != 0 {
					leftIndex := stack[len(stack)-1]
					rightIndex := i

					h := min(height[leftIndex], height[rightIndex]) - height[buttonIndex]
					w := rightIndex - leftIndex - 1

					area := h * w
					sum += area
				}
			}

			stack = append(stack, i)
		}
	}
	return sum
}
