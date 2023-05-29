package meeting

type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

type ListNode struct {
	Val  int       `json:"val"`
	Next *TreeNode `json:"next"`
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
