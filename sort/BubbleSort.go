/*
*  @author: didichuxing.com
 */

package sort

/*
	冒泡排序
	时间复杂度是O(n^2), 最坏时间复杂度O(n^2), 稳定，因为不会该表任意的时间
*/

func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j+1] < nums[j] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}
