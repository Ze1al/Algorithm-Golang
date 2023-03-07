package greed

// 跳跃游戏二
// https://leetcode.cn/problems/jump-game-ii/
func jump(nums []int) int {
	var res = 0
	curRange := 0
	nextRange := 0 // 代表走走完当前这一步，nextRange能覆盖的最大距离，如果覆盖面能到数组末尾，则为到底

	for i := 0; i < len(nums); i++ {
		nextRange = max(nextRange, nums[i]+i) // 当前这一步能覆盖的最大范围
		if i == curRange {                    // 到了这一步不得不走了
			if curRange < len(nums)-1 { // 还是没走到目的地
				res += 1
				curRange = nextRange
				if nextRange >= len(nums)-1 {
					break
				}
			} else {
				break
			}
		}
	}
	return res
}
