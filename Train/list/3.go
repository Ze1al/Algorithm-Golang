package list

//https://leetcode.cn/problems/squares-of-a-sorted-array/
//有序数组的平方
func sortedSquares(nums []int) []int {
	l := 0
	r := len(nums) - 1

	res := make([]int, len(nums))
	index := len(nums) - 1

	for l <= r {
		if nums[l]*nums[l] < nums[r]*nums[r] {
			res[index] = nums[r] * nums[r]
			index--
			r--
		} else {
			res[index] = nums[l] * nums[l]
			index--
			l++
		}
	}
	return res
}
