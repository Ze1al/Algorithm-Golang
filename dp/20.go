package dp

// 打家劫舍
// https://leetcode.cn/problems/house-robber/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	// 价值： nums[i]
	// 背包：length
	// 先遍历背包，在遍历价值
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}
