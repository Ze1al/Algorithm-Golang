package meeting

import (
	"strconv"
	"strings"
)

// 猿辅导
// 有效IP
var res []string

func restoreIpAddresses(s string) []string {
	res = []string{}
	backTravel(s, []string{}, 0)
	return res
}

func backTravel(s string, path []string, startIndex int) {
	if startIndex == len(s) && len(path) == 4 {
		ip := strings.Join(path, ".")
		res = append(res, ip)
		return
	}
	if len(path) > 4 {
		return
	}
	for i := startIndex; i < len(s); i++ {
		sub := s[startIndex : i+1]
		if isValidIp(sub) {
			path = append(path, sub)
			backTravel(s, path, i+1)
			path = path[:len(path)-1]
		} else {
			continue
		}
	}
}

func isValidIp(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if num > 255 {
		return false
	}
	return true
}
