package offer

// 插入位置
// 其实就是找到第一个大于target的索引
// 左闭右开区间
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
