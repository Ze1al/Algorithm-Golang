package dp

// 零钱兑换
// https://leetcode.cn/problems/coin-change-ii/
func change(amount int, coins []int) int {
	// 1. dp[j]代表总价值为j的，一共有多少种不同组合
	// amount就是背包，coins就是物品
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}
