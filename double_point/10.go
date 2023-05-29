package double_point

import "sort"

// 三数之和
// 满足 num[i]+num[j]+num[k] == 0
// 用双指针的想法，
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res = [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right {
					if nums[left] == nums[left+1] {
						left += 1
					} else {
						break
					}
				}
				for left < right {
					if nums[right] == nums[right-1] {
						right -= 1
					} else {
						break
					}
				}
				left += 1
				right -= 1
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return res
}
