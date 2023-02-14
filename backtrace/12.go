package backTrace

import "sort"

var res12 [][]int

func permuteUnique(nums []int) [][]int {
	res12 = [][]int{}
	used := make([]bool, len(nums))
	sort.Ints(nums)

	backTracking12(nums, []int{}, used)
	return res12
}

func backTracking12(nums []int, path []int, used []bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		res12 = append(res12, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}

		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			backTracking12(nums, path, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
}
