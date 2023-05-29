package offer

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt64
	sum := 0 // 代表当前窗口中的sum
	i := 0

	for j := 0; j < len(nums); j++ {
		sum += nums[j]

		for sum >= target {
			temp := j - i + 1
			if temp < res {
				res = temp
			}
			sum -= nums[i]
			i++
		}
	}

	if res == math.MaxInt64 {
		return 0
	}
	return res
}
