package dp

// 单词拆分
// https://leetcode.cn/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	// 1. 确定含义
	// dp[j] 代表字符串长度为i的，dp[j] = true 代表存在
	dp := make([]bool, len(s)+1)

	// 2. 初始化
	dp[0] = true

	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}

	// 3.递推公式
	// 由于是排列问题，所以先遍历背包，再遍历物品
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
