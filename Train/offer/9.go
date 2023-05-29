package offer

// 值和下标之差都在给定范围之内
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for left := 0; left < len(nums); left += 1 {
		for right := left + 1; right <= left+k && right < len(nums); right += 1 {
			if abs(nums[left]-nums[right]) <= t {
				return true
			}
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
