package offer

// 最小路径和
// 典型的DP

// dp[i][j] 代表走到下标为i和j位置需要的最小步数
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < len(dp); i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
