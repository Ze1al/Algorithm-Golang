package monotonic_stack

// 下一个更大元素一
// 首先计算出nums2中各个元素的右边最大值
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{} // 单调递增，存放下标
	nextGreaterNum := make([]int, len(nums2))

	for i := 0; i < len(nextGreaterNum); i++ {
		nextGreaterNum[i] = -1
	}

	for i := 0; i < len(nums2); i++ {
		index := len(stack) - 1
		for index >= 0 && nums2[stack[index]] <= nums2[i] {
			val := stack[len(stack)-1]
			nextGreaterNum[val] = nums2[i]

			stack = stack[:len(stack)-1]
			index -= 1
		}

		stack = append(stack, i)
	}

	var indexMap = make(map[int]int)
	for i, v := range nums2 {
		indexMap[v] = nextGreaterNum[i]
	}

	var res = make([]int, len(nums1))
	for i, v := range nums1 {
		nextGreaterNum, ok := indexMap[v]
		if !ok {
			continue
		}
		res[i] = nextGreaterNum
	}
	return res
}
