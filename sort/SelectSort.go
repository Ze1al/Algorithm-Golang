/*
*  @author: didichuxing.com
 */

package sort

// 选择排序，每次选择出最小的数值，然后放到第一位， 不稳定

func SelectSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}
