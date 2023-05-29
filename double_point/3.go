package double_point

// 移动零
// 还是双指针的思想
// 如果nums[i] != 0 则交换到前面
// 如果等于0 则不移动
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow += 1
		}
		fast += 1
	}
	return
}
