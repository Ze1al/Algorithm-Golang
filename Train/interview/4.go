package interview

import "sort"

// 数字n， 大小为m的数组，三个数字之和

// 方法一：回溯
// 方法二：双指针

// 返回具体满足条件的情况
// 返回具体的下标
func threeSum(nums []int, n int) [][]int {
	// 第一步，排序，方便后序减枝
	sort.Ints(nums)

	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	var res [][]int

	for i := 0; i < len(nums); i++ {
		// 因为是有序，所以如果nums[i]大于n，则肯定后续不会满足
		if nums[i] >= n {
			break
		}

		left := i + 1
		for left < len(nums) {
			target := n - nums[left] - nums[i]

			index, ok := m[target]
			if ok && index > left {
				res = append(res, []int{i, left, index})
				left += 1
			} else {
				left += 1
			}
		}
	}
	return res
}
