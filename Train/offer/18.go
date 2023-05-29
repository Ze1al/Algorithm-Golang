package offer

// 通过二分法，查找第一个元素和第二个元素
func searchRange(nums []int, target int) []int {
	left := searchLeft(nums, target)
	right := searchRight(nums, target)

	if left == -2 || right == -2 {
		return []int{-1, -1}
	}
	if right-left > 1 {
		return []int{left + 1, right - 1}
	}

	return []int{-1, -1}
}

// 获取左边界
// 获取到相等的时候，应该左边走
func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1

	ans := -2

	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] >= target {
			right = mid - 1
			ans = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

// 找到右边界，获取到相等的时候，应该往右边走
func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1

	ans := -2

	for left <= right {
		mid := (right-left)/2 + left
		if target >= nums[mid] {
			left = mid + 1
			ans = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
