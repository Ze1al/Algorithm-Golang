package dp

// https://leetcode.cn/problems/integer-break/
// 整数拆分
// 1、dp[i] 代表i的最大拆解乘积
// 2、dp[i] = j * dp[i-j]   或者   dp[i] = j * dp[i-j]      或者       dp[i]
// 3、初始化 dp[2] = 1
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = threeMax(dp[i], j*(i-j), j*dp[i-j])
		}
	}
	return dp[n]
}

func threeMax(x, y, z int) int {
	max := x
	if max < y {
		max = y
	}
	if max < z {
		max = z
	}
	return max
}
