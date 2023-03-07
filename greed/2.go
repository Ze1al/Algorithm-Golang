package greed

// 摆动序列
// https://leetcode.cn/problems/wiggle-subsequence/
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return ans
}
