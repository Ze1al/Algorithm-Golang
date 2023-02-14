package backTrace

import "sort"

// 找子集
// https://leetcode.cn/problems/subsets-ii/

var res9 [][]int

func subsetsWithDup(nums []int) [][]int {
	res9 = [][]int{}
	path := []int{}
	used := make([]bool, len(nums))

	sort.Ints(nums)

	backTracking9(nums, path, used, 0)
	return res9
}

func backTracking9(nums []int, path []int, used []bool, startIndex int) {
	temp := make([]int, len(path))
	copy(temp, path)
	res9 = append(res9, temp)

	for i := startIndex; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}

		path = append(path, nums[i])
		used[i] = true
		backTracking9(nums, path, used, i+1)
		used[i] = false
		path = path[:len(path)-1]
	}
}
