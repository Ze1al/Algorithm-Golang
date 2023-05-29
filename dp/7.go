package dp

// https://leetcode.cn/problems/unique-binary-search-trees/
// 不同的二叉搜索树个数
// 1: dp[i] 代表i个节点有多少种不同组合方式
// 2: dp[i] += dp[j-1] * dp[i-j]
// 3: dp[0] = 1, dp[1] = 1
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
