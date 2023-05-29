package offer

// 移除相同值的元素
func removeElement(nums []int, val int) int {
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow += 1
		}
		fast += 1
	}
	return slow
}
