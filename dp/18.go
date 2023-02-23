package dp

import (
	"fmt"
	"math"
)

// 完全平方和
// https://leetcode.cn/problems/perfect-squares/
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 1; i < n+1; i++ {
		dp[i] = math.MaxInt64
	}

	for i := 1; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
		fmt.Println(dp[i])
	}

	if dp[n] == math.MaxInt64 {
		return -1
	}
	return dp[n]

}
