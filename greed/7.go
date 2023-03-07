package greed

import (
	"math"
	"sort"
)

// K 次取反后最大化的数组和
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) < math.Abs(float64(nums[j]))
	})

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < 0 && k > 0 {
			nums[i] = -1 * nums[i]
			k -= 1
		}
	}

	if k%2 == 1 {
		nums[0] = -1 * nums[0]
	}

	var res int
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}
