package dp

/// https://leetcode.cn/problems/combination-sum-iv
// 组合总和
// 如果求组合数就是外层for循环遍历物品，内层for遍历背包。
// 如果求排列数就是外层for遍历背包，内层for循环遍历物品。

func combinationSum4(nums []int, target int) int {
	//定义dp数组
	dp := make([]int, target+1)
	// 初始化
	dp[0] = 1
	// 遍历顺序, 先遍历背包,再循环遍历物品
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
