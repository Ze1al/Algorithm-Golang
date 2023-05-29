package double_point

import "strings"

// 反转字符串中的单词
func reverseWords(s string) string {
	arr := strings.Split(s, " ")

	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left += 1
		right -= 1
	}

	return strings.TrimSpace(strings.Join(arr, " "))
}
