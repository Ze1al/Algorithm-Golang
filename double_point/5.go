package double_point

// 有序数组的平方
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i, j := 0, len(nums)-1

	index := len(res) - 1

	for i <= j {
		if nums[i]*nums[i] < nums[j]*nums[j] {
			res[index] = nums[j] * nums[j]
			j -= 1
		} else {
			res[index] = nums[i] * nums[i]
			i += 1
		}

		index -= 1
	}
	return res
}
