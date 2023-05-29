package backTrace

// 是否为回文字符串
func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		} else {
			i += 1
			j -= 1
		}
	}
	return true
}

var res6 [][]string

// N种切分方式，保证其中的子串都是回文字符串
func partition(s string) [][]string {
	res6 = [][]string{}

	backTracking6(s, []string{}, 0)
	return res6
}

func backTracking6(s string, path []string, startIndex int) {
	if startIndex == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		res6 = append(res6, temp)
		return
	}

	for i := startIndex; i < len(s); i++ {
		s2 := s[startIndex : i+1] // 截取
		if !isPalindrome(s2) {
			continue
		} else {
			// 如果是回文字符串，则添加到path路径中
			path = append(path, s2)
			backTracking6(s, path, i+1)
			path = path[:len(path)-1]
		}
	}
}
