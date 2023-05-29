/*
*  @author: didichuxing.com
 */

package sort

// 插入排序 时间复杂度O(n^2)，稳定

func InsertSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if nums[j] > nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}
