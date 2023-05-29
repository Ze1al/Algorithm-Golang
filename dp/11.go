package dp

// https://leetcode.cn/problems/last-stone-weight-ii/
// 最后一块石头的重量

func lastStoneWeightII(stones []int) int {
	dp := make([]int, 15001)
	sum := 0
	for _, v := range stones {
		sum += v
	}

	target := sum / 2
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - dp[target] - dp[target]
}
