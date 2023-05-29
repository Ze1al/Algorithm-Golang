package offer

// 排列的数目
// 顺序可以不限制，因此遍历是先背包，再物品
// 无限次，所以是完全背包

// dp[j] 代表凑够j的方式

// dp[j] += dp[j-nums[i]]

// dp[0] = 1

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for j := 1; j < target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
