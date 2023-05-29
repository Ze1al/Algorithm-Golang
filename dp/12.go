package dp

import "math"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if abs(target) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}

	// dp[j] 为填满J的包存在几种情况
	bag := (sum + target) / 2
	dp := make([]int, bag+1)

	// dp[0] = 1初始化为1
	dp[0] = 1

	// 遍历
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[bag]
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
