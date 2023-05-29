package offer

// left - right = target
// left + right = sum
// left = (sum + target) / 2

// dp[j] 代表凑够j 的最大价值，0-1背包

// dp[j] = max(dp[j], dp[j-nums[i]] + nums[i])

// dp[0] = 0

func findTargetSumWays(nums []int, target int) int {
	sum := func(nums []int) int {
		if len(nums) == 0 {
			return 0
		}
		s := nums[0]
		for i := 1; i < len(nums); i++ {
			s += nums[i]
		}
		return s
	}

	sum2 := sum(nums)
	if target > sum2 {
		return 0
	}
	if (target+sum2)%2 == 1 {
		return 0
	}
	left := (sum2 + target) / 2

	dp := make([]int, left+1)
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := left; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[left]
}
