package greed

import "math"

// 最大子数组和
// 贪心，保证局部是最小的
func maxSubArray(nums []int) int {
	var (
		res = math.MinInt64
		sum = 0
	)

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > res {
			res = sum
		}

		if sum <= 0 {
			sum = 0
		}
	}
	return res
}
