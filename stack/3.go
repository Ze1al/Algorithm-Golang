package stack

var m = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

// 有效括号
func isValid(s string) bool {
	stack := []string{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, string(s[i]))
			continue
		}

		left, ok := m[string(s[i])]
		if ok && len(stack) > 0 && stack[len(stack)-1] == left {
			stack = stack[:len(stack)-1]
			continue
		} else {
			return false
		}
	}
	return len(stack) == 0
}
