package dp

// https://leetcode.cn/problems/ones-and-zeroes/
// 一和零
func findMaxForm(strs []string, m int, n int) int {
	// 1. dp[i][j] 代表i个0和j个1的最长子集
	// m 代表0的个数 n代表1的个数
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	// 2. 初始化

	// 3. 递推公式
	// dp[i][j] = max(dp[i-1][j], dp[i-zero_num][j-one_num] + 1)
	for _, v := range strs {
		oneNum, zeroNum := 0, 0
		for i := 0; i < len(v); i++ {
			if v[i] == '0' {
				zeroNum += 1
			} else {
				oneNum += 1
			}
		}

		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}
