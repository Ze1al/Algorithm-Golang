package offer

// 分割等和子集

// 背包问题，0-1背包
// target是背包的容量

// dp[j] 代表累加等于j 的 最大价值

// dp[j] = max(dp[j], dp[j-nums[i]] + nums[i])
// dp[0] = 0

func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	dp := make([]int, target+1)

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
