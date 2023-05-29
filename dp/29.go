package dp

import "fmt"

// 买卖股票含有手续费
func maxProfit6(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -1 * prices[0] // 持有股票
	dp[0][1] = 0              // 不持有股票

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	fmt.Printf("dp=%v", dp)

	return dp[len(prices)-1][1]
}
