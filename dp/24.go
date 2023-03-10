package dp

// 买卖股票的最大收益
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	dp := make([][]int, 2)
	for i, _ := range dp {
		dp[i] = make([]int, len(prices))
	}

	// dp[i][0]代表第i天持有股票的最高所得
	// dp[i][1]代表第i天不持有股票的最高所得

	dp[0][0] -= prices[0]
	dp[0][1] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}
