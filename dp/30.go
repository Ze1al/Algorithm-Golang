package dp

// 最长上升子序列
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)) // dp[i] 代表以i结尾的，最大长度
	for i, _ := range dp {
		dp[i] = 1
	}

	res := 0

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
