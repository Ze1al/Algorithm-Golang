package greed

// 跳跃游戏
// 贪心，判断获取最大范围能不能覆盖终点
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return nums[0] > 0
	}
	rang := nums[0]
	for i := 1; i < len(nums) && i <= rang; i++ {
		rang = max(rang, nums[i]+i)
	}

	return rang >= len(nums)-1
}
