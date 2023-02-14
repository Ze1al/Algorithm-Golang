package stack

import "strings"

// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
func removeDuplicates(s string) string {
	stack := []string{}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || (len(stack) > 0 && stack[len(stack)-1] != string(s[i])) {
			stack = append(stack, string(s[i]))
			continue
		}
		if len(stack) > 0 && stack[len(stack)-1] == string(s[i]) {
			stack = stack[:len(stack)-1]
			continue
		}
	}
	return strings.Join(stack, "")
}
