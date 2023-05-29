package backTrace

// 递增子序列
// https://leetcode.cn/problems/non-decreasing-subsequences/

var res10 [][]int

func findSubsequences(nums []int) [][]int {
	res10 = [][]int{}

	backTracking10(nums, []int{}, 0)
	return res10
}

func backTracking10(nums []int, path []int, startIndex int) {
	if len(path) > 1 {
		temp := make([]int, len(path))
		copy(temp, path)
		res10 = append(res10, temp)
	}

	used := make(map[int]bool)

	for i := startIndex; i < len(nums); i++ {
		if used[nums[i]] {
			continue
		}
		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			path = append(path, nums[i])
			used[nums[i]] = true
			backTracking10(nums, path, i+1)
			path = path[:len(path)-1]
		}
	}
}
