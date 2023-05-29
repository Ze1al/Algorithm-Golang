package meeting

import "math"

// 快手
// 兑换零钱
// 零钱是无限个数的，所以是完全背包
// dp[j] 凑够j元，需要的最少个数
// 背包：amount
// 物品：coins
// 价值：coins
// 递推公式：dp[j] = min(dp[j], dp[i-coins[j]] + 1)
// 初始化：dp[0] = 0

// 递推顺序：完全背包，先遍历物品，在遍历背包，正序
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 需要赋值一下，避免被0覆盖
	for j := 1; j < len(dp); j++ {
		dp[j] = math.MaxInt
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]]+1 == math.MaxInt {
				continue
			}
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
