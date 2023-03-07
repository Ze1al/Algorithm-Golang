package monotonic_stack

// 下一个更大的元素
// https://leetcode.cn/problems/next-greater-element-i/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	res := make([]int, len(nums1))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}

	// 从value映射到index
	value2IndexMap := make(map[int]int)
	for i := 0; i < len(res); i++ {
		value2IndexMap[nums1[i]] = i
	}

	stack = append(stack, 0)

	for i := 1; i < len(nums2); i++ {
		right := stack[len(stack)-1]
		if nums2[right] > nums2[i] {
			stack = append(stack, i)
		}
		if nums2[right] == nums2[i] {
			stack = append(stack, i)
		}

		if nums2[right] < nums2[i] {
			for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] {
				if index, ok := value2IndexMap[nums2[stack[len(stack)-1]]]; ok {
					res[index] = nums2[i]
				}
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, i)
		}
	}

	return res
}
