package offer

// 有序数组的平方
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1

	res := make([]int, len(nums))
	k := len(res) - 1
	for left < right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			res[k] = nums[right] * nums[right]
			right -= 1
		} else {
			res[k] = nums[left] * nums[left]
			left += 1
		}
		k -= 1
	}
	return res
}
