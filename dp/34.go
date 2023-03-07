package dp

// 最大子数组和
// https://leetcode.cn/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	var res = nums[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
