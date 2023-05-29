package double_point

// 移除元素，相等的元素移除，要求原地
// 返回移除以后得下标
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
