package offer

// 有效的完全平方数
func isPerfectSquare(num int) bool {
	left := 0
	right := num
	for left <= right {
		mid := (right-left)/2 + left
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
