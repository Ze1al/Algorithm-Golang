package double_point

// 替换空格
func replaceSpace(s string) string {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count += 1
		}
	}

	res := make([]byte, len(s)+count*2)

	i, j := len(s)-1, len(res)-1
	for i >= 0 && j >= 0 {
		if s[i] != ' ' {
			res[j] = s[i]
			j -= 1
			i -= 1
		} else {
			res[j] = '0'
			res[j-1] = '2'
			res[j-2] = '%'

			j -= 3
			i -= 1
		}
	}
	return string(res)
}
