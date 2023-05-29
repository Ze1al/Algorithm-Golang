package double_point

// 删除有序数组中的重复项目
// 快慢指针经典题目
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	fast, slow := 1, 1
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow += 1
		}
		fast += 1
	}
	return slow
}
