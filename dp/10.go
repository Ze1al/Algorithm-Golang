package dp

import "sort"

// 分割等和子集
// https://leetcode.cn/problems/partition-equal-subset-sum/

// 属于 0-1背包的变形
func canPartition(nums []int) bool {
	target := sum(nums) / 2
	sort.Ints(nums)

	if sum(nums)%2 == 1 {
		return false
	}

	// 1、dp[i][j]代表下标为 [0,i]任取元素，target == j的最大价值
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
	}

	// 2、初始化
	for i := 0; i < len(dp[0]); i++ {
		if i >= nums[0] {
			dp[0][i] = nums[0]
		}
	}

	// 3、递推
	for i := 1; i < len(dp); i++ {
		for j := 1; j <= target; j++ {
			// 说明装不下nums[i]
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			}
			if target == dp[i][j] {
				return true
			}
		}
	}
	return false
}

func canPartition2(nums []int) bool {
	target := sum(nums) / 2
	sort.Ints(nums)

	if sum(nums)%2 == 1 {
		return false
	}

	// dp[j] 代表target为J的nums和
	dp := make([]int, target+1)

	for i := 0; i < len(nums); i++ {
		for j := target; j > 0; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}

func sum(nums []int) int {
	target := 0
	for _, v := range nums {
		target += v
	}
	return target
}
