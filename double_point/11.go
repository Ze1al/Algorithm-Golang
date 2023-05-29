package double_point

import (
	"sort"
)

// 四数之和
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > target && nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := len(nums) - 1

			sum := nums[i] + nums[j]

			for left < right {
				if nums[left]+nums[right] == target-sum {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left += 1
					}
					for left < right && nums[right] == nums[right-1] {
						right -= 1
					}
					left += 1
					right -= 1
				} else if nums[left]+nums[right] < target-sum {
					left += 1
				} else {
					right -= 1
				}
			}
		}
	}

	return res
}
