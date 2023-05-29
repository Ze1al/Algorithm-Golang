package double_point

// 比较含退格的字符串
// 方法一，使用栈
func backspaceCompare1(s string, t string) bool {
	return build(s) == build(t)
}

func build(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '#' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
		if s[i] != '#' {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func backspaceCompare(s, t string) bool {
	i, j := len(s)-1, len(t)-1
	skipS, skipT := 0, 0

	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS += 1
				i -= 1
			} else if skipS > 0 {
				skipS -= 1
				i -= 1
			} else {
				break
			}
		}

		for j >= 0 {
			if t[j] == '#' {
				skipT += 1
				j -= 1
			} else if skipT > 0 {
				skipT -= 1
				j -= 1
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else {
			if i >= 0 || j >= 0 {
				return false
			}
		}

		i -= 1
		j -= 1
	}
	return true
}
