package offer

import "math"

// 硬币数量是无限的，所以代表的是完全背包

// dp[j] 代表凑够j金额，需要的最少的硬币数目
// dp[j] = min(dp[j], dp[j-nums[i]])
// dp[0] = 1

// 遍历顺序，先遍历物品，再遍历背包

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < amount+1; j++ {
			if dp[j-coins[i]] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
